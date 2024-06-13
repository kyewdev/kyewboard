package util

import "strconv"


func TasklblInput(index int) string {
    return "tasklblInput-" + strconv.Itoa(index)
}

func TasklblInputHash(index int) string {
    return "#tasklblInput-" + strconv.Itoa(index)
}
