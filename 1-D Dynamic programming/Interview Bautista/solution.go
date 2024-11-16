package main

import "fmt"

func maximizeItems(itemsFloat map[string]float64, amount float64) ([]string, float64) {
	// Quitar decimales y trabajar con enteros
	amountInt := int(amount * 100)
	items := make(map[string]int, len(itemsFloat))
	for id, price := range itemsFloat {
		items[id] = int(price * 100)
	}

	n := len(items)
	dp := make([]map[int]bool, n+1)
	itemsIndex := make([]string, len(items))
	for i := range dp {
		dp[i] = make(map[int]bool)
	}

	dp[0][0] = true // Caso base
	i := 0
	for id, itemPrice := range items {
		itemsIndex[i] = id
		for prevAmount := range dp[i] {
			dp[i+1][prevAmount] = true
			newAmount := prevAmount + itemPrice
			if newAmount <= amountInt {
				dp[i+1][newAmount] = true
			}
		}
		i++
	}

	// Encontrar el monto máximo alcanzable en la tabulación
	var selectedTotal int
	for price := range dp[n] {
		if price <= amountInt && price > selectedTotal {
			selectedTotal = price
		}
	}
	total := float64(selectedTotal) / 100.0

	// Reconstruir la lista de ítems seleccionados
	var selected []string
	for i := n; i > 0 && selectedTotal > 0; i-- {
		itemID := itemsIndex[i-1]
		itemPrice := items[itemID]

		if used := dp[i-1][selectedTotal-itemPrice]; used {
			selected = append(selected, itemID)
			selectedTotal -= itemPrice
		}
	}

	return selected, total
}

func main() {
	itemsData := map[string]float64{
		"MLA1": 100.0,
		"MLA2": 210.0,
		"MLA3": 260.0,
		"MLA4": 80.0,
		"MLA5": 90.0,
	}

	fmt.Println(maximizeItems(itemsData, 500))
}
