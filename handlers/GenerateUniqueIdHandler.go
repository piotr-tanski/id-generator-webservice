package handlers

import (
	"id_generator/generator"
	"net/http"
)

type GenerateUniqueIdHandler struct {
	Generator *generator.IdGenerator
}

func (handler *GenerateUniqueIdHandler) GenerateId() (int, map[string]any) {
	id, timestamp, machine_id, seqno := handler.Generator.Generate()
	return http.StatusOK, toMap(id, timestamp, machine_id, seqno)
}

func toMap(id int64, timestamp int64, machine_id int, seqno int) map[string]any {
	result := make(map[string]any)
	result["id"] = id
	result["timestamp"] = timestamp
	result["machine_id"] = machine_id
	result["seqno"] = seqno
	return result
}
