package testutil

import (
	"context"

	"github.com/CPunch/gopenfusion/internal/db"
	"github.com/CPunch/gopenfusion/internal/redis"
	"github.com/alicebob/miniredis/v2"
	"github.com/bitcomplete/sqltestutil"
)

// SetupEnvironment spawns a postgres container and returns a db and redis handler
// along with a cleanup function
func SetupEnvironment(ctx context.Context) (*db.DBHandler, *redis.RedisHandler, func()) {
	// spawn postgres container
	psql, err := sqltestutil.StartPostgresContainer(ctx, "15")
	if err != nil {
		panic(err)
	}

	// open db handler
	testDB, err := db.OpenFromConnectionString("postgres", psql.ConnectionString()+"?sslmode=disable")
	if err != nil {
		psql.Shutdown(ctx)
		panic(err)
	}

	if err = testDB.Setup(); err != nil {
		psql.Shutdown(ctx)
		panic(err)
	}

	// start miniredis
	r, err := miniredis.Run()
	if err != nil {
		psql.Shutdown(ctx)
		panic(err)
	}

	// open redis handler
	rh, err := redis.OpenRedis(r.Addr())
	if err != nil {
		psql.Shutdown(ctx)
		panic(err)
	}

	return testDB, rh, func() {
		psql.Shutdown(ctx)
		rh.Close()
		r.Close()
	}
}
