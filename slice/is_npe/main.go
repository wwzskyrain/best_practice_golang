package main

import "log"

func main() {
	finalOaIds := make([]string, 0)
	oaIDs := getOneEmptyEleOaIds()
	log.Printf("oaIds = %+v len=%d", oaIDs, len(oaIDs))
	for _, oaId := range oaIDs {
		if len(oaId) > 0 {
			finalOaIds = append(finalOaIds, oaId)
		}
	}
	log.Printf("finalOaIds = %+v, len=%d", finalOaIds, len(finalOaIds))

}

func getNilOaIds() []string {
	return nil
}

func getEmptyEleOaIds() []string {
	return []string{"", "", "1"}
}

func getOneEmptyEleOaIds() []string {
	return []string{""}
}
