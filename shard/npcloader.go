package shard

import (
	"encoding/json"
	"log"
	"os"

	"github.com/CPunch/gopenfusion/internal/config"
	"github.com/CPunch/gopenfusion/shard/entity"
)

type NPCData struct {
	NPCs map[string]entity.NPC `json:"NPCs"`
}

func (server *ShardServer) LoadNPCs() {
	log.Print("Loading NPCs...")

	data, err := os.ReadFile(config.GetTDataPath() + "/NPCs.json")
	if err != nil {
		panic(err)
	}

	// yes, we have to do it this way so our NPCs IDs will be incremented and unique
	var NPCs NPCData
	json.Unmarshal(data, &NPCs)
	for _, npc := range NPCs.NPCs {
		server.addEntity(entity.NewNPC(npc.X, npc.Y, npc.Z, npc.Angle, npc.NPCType))
	}

	log.Printf("Loaded %d NPCs!", len(NPCs.NPCs))
}
