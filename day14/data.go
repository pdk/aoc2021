package main

var start = "OFSNKKHCBSNKBKFFCVNB"

type pair struct {
	left, right rune
}

type ruleMap map[pair]rune

var rules = ruleMap{
	{'K', 'C'}: 'F',
	{'C', 'O'}: 'S',
	{'F', 'H'}: 'K',
	{'V', 'P'}: 'P',
	{'K', 'F'}: 'S',
	{'S', 'V'}: 'O',
	{'C', 'B'}: 'H',
	{'P', 'N'}: 'F',
	{'N', 'C'}: 'N',
	{'B', 'C'}: 'F',
	{'N', 'P'}: 'O',
	{'S', 'K'}: 'F',
	{'H', 'S'}: 'C',
	{'S', 'N'}: 'V',
	{'O', 'P'}: 'F',
	{'O', 'N'}: 'N',
	{'F', 'K'}: 'N',
	{'S', 'H'}: 'B',
	{'H', 'N'}: 'N',
	{'B', 'O'}: 'V',
	{'V', 'K'}: 'H',
	{'S', 'C'}: 'K',
	{'K', 'P'}: 'O',
	{'V', 'O'}: 'V',
	{'H', 'C'}: 'P',
	{'B', 'K'}: 'B',
	{'V', 'H'}: 'N',
	{'P', 'V'}: 'O',
	{'H', 'B'}: 'H',
	{'V', 'S'}: 'F',
	{'K', 'K'}: 'B',
	{'H', 'H'}: 'B',
	{'C', 'F'}: 'F',
	{'P', 'H'}: 'C',
	{'N', 'S'}: 'V',
	{'S', 'O'}: 'P',
	{'N', 'V'}: 'K',
	{'B', 'P'}: 'N',
	{'S', 'F'}: 'V',
	{'S', 'S'}: 'K',
	{'F', 'P'}: 'N',
	{'P', 'C'}: 'S',
	{'O', 'H'}: 'B',
	{'C', 'H'}: 'H',
	{'V', 'V'}: 'S',
	{'V', 'N'}: 'O',
	{'O', 'B'}: 'K',
	{'P', 'F'}: 'H',
	{'C', 'S'}: 'C',
	{'P', 'P'}: 'O',
	{'N', 'F'}: 'H',
	{'S', 'P'}: 'P',
	{'O', 'S'}: 'V',
	{'B', 'B'}: 'P',
	{'N', 'O'}: 'F',
	{'V', 'B'}: 'V',
	{'H', 'K'}: 'C',
	{'N', 'K'}: 'O',
	{'H', 'P'}: 'B',
	{'H', 'V'}: 'V',
	{'B', 'F'}: 'V',
	{'K', 'O'}: 'F',
	{'B', 'V'}: 'H',
	{'K', 'V'}: 'B',
	{'O', 'F'}: 'V',
	{'N', 'B'}: 'F',
	{'V', 'F'}: 'C',
	{'P', 'B'}: 'B',
	{'F', 'F'}: 'H',
	{'C', 'P'}: 'C',
	{'K', 'H'}: 'H',
	{'N', 'H'}: 'P',
	{'P', 'S'}: 'P',
	{'P', 'K'}: 'P',
	{'C', 'C'}: 'K',
	{'B', 'S'}: 'V',
	{'S', 'B'}: 'K',
	{'O', 'O'}: 'B',
	{'O', 'K'}: 'F',
	{'B', 'H'}: 'B',
	{'C', 'V'}: 'F',
	{'F', 'N'}: 'V',
	{'C', 'N'}: 'P',
	{'K', 'B'}: 'B',
	{'F', 'O'}: 'H',
	{'P', 'O'}: 'S',
	{'H', 'O'}: 'H',
	{'C', 'K'}: 'B',
	{'K', 'N'}: 'C',
	{'F', 'S'}: 'K',
	{'O', 'C'}: 'P',
	{'F', 'V'}: 'N',
	{'O', 'V'}: 'K',
	{'B', 'N'}: 'H',
	{'H', 'F'}: 'V',
	{'V', 'C'}: 'S',
	{'F', 'B'}: 'S',
	{'N', 'N'}: 'P',
	{'F', 'C'}: 'B',
	{'K', 'S'}: 'N',
}
