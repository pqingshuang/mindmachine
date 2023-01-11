package shares

func GetAllExpenses() (s map[string]Share) {
	currentState.mutex.Lock()
	defer currentState.mutex.Unlock()
	//for _, share := range currentState.data {
	//	s = append(s, share)
	//}
	return currentState.data
}
