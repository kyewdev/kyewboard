package util

import (
	"strconv"

)


func TasklblInput(questindex int, objectiveindex int) string {
    return "tasklblInput-" + strconv.Itoa(questindex) + "-" + strconv.Itoa(objectiveindex)
}

func TasklblInputHash(questindex int, objectiveindex int) string {
    return "#tasklblInput-" + strconv.Itoa(questindex) + "-" + strconv.Itoa(objectiveindex)
}
