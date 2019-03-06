package agent

import (
    "gopkg.in/mgo.v2/bson"
)

// Config - agent JSON config
type Config struct {
    ID          bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
    AgentID     string 	        `json:"agentId" bson:"agentId,omitempty"`
    TmpAdjust   string          `json:"temperature" bson:"tmpAdjustment"`
}
