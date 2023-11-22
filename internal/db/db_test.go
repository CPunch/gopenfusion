package db

import (
	"context"
	"os"
	"testing"

	"github.com/CPunch/gopenfusion/internal/protocol"
	"github.com/bitcomplete/sqltestutil"
)

var (
	testDB *DBHandler
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	psql, err := sqltestutil.StartPostgresContainer(ctx, "15")
	if err != nil {
		panic(err)
	}
	defer psql.Shutdown(ctx)

	testDB, err = OpenFromConnectionString("postgres", psql.ConnectionString()+"?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer testDB.Close()

	if err := testDB.Setup(); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestDBAccount(t *testing.T) {
	if _, err := testDB.NewAccount("test", "test"); err != nil {
		t.Error(err)
	}

	// now try to retrieve account data
	acc, err := testDB.TryLogin("test", "test")
	if err != nil {
		t.Error(err)
	}

	if acc.Login != "test" {
		t.Error("account username is not test")
	}

	if _, err = testDB.TryLogin("test", "wrongpassword"); err != ErrLoginInvalidPassword {
		t.Error("expected ErrLoginInvalidPassword")
	}
}

/*

test data has been collected by running the following commands in a postgresql shell started with:
docker exec -it gofusion-postgresql-1 psql -U gopenfusion -W gopenfusion

gopenfusion=# SELECT * FROM Players;
 playerid | accountid | firstname | lastname | namecheck | slot | created | lastlogin | level | nano1 | nano2 | nano3 | appearanceflag | tutorialflag | payzoneflag | xcoordinate | ycoordinate | zcoordinate | angle |  hp  | fusionmatter | taros | batteryw | batteryn | mentor | currentmissionid | warplocationflag |         skywaylocationflag         |            firstuseflag            |                                                                                                                               quests
----------+-----------+-----------+----------+-----------+------+---------+-----------+-------+-------+-------+-------+----------------+--------------+-------------+-------------+-------------+-------------+-------+------+--------------+-------+----------+----------+--------+------------------+------------------+------------------------------------+------------------------------------+--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
        1 |         1 | Neil      | Mcscout  |         1 |    1 |   76476 |     76476 |     1 |     0 |     0 |     0 |              1 |            1 |           0 |      632032 |      187177 |       -5500 |     0 | 1000 |            0 |     0 |        0 |        0 |      5 |                0 |                0 | \x00000000000000000000000000000000 | \x00000000000000000000000000000000 | \x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000

gopenfusion=# SELECT * FROM Appearances;
 playerid | body | eyecolor | facestyle | gender | haircolor | hairstyle | height | skincolor
----------+------+----------+-----------+--------+-----------+-----------+--------+-----------
        1 |    2 |        3 |         5 |      1 |        11 |         3 |      0 |         9

gopenfusion=# SELECT * FROM Inventory;
 playerid | slot | id  | type | opt | timelimit
----------+------+-----+------+-----+-----------
        1 |    1 | 328 |    1 |   1 |         0
        1 |    2 | 359 |    2 |   1 |         0
        1 |    3 | 194 |    3 |   1 |         0

*/

func TestDBPlayer(t *testing.T) {
	if _, err := testDB.NewAccount("testplayer", "test"); err != nil {
		t.Error(err)
	}

	// now try to retrieve account data
	acc, err := testDB.TryLogin("testplayer", "test")
	if err != nil {
		t.Error(err)
	}

	plrID, err := testDB.NewPlayer(acc.AccountID, "Neil", "Mcscout", 1)
	if err != nil {
		t.Error(err)
	}

	if err = testDB.FinishPlayer(&protocol.SP_CL2LS_REQ_CHAR_CREATE{
		PCStyle: protocol.SPCStyle{
			IPC_UID:     int64(plrID),
			INameCheck:  1,
			SzFirstName: "Neil",
			SzLastName:  "Mcscout",
			IGender:     1,
			IFaceStyle:  5,
			IHairColor:  11,
			ISkinColor:  9,
			IEyeColor:   3,
			IHeight:     0,
			IBody:       2,
			IHairStyle:  3,
		},
		SOn_Item: protocol.SOnItem{
			IEquipUBID:   328,
			IEquipLBID:   359,
			IEquipFootID: 194,
		},
	}, acc.AccountID); err != nil {
		t.Error(err)
	}

	if err = testDB.FinishTutorial(plrID, acc.AccountID); err != nil {
		t.Error(err)
	}
}