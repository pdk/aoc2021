package main

var testData = []*node{
	np(1, 1),
	np(2, 2),
	np(3, 3),
	np(4, 4),
	np(5, 5),
	np(6, 6),
}

var largerExample = []*node{

	np(np(np(0, np(4, 5)), np(0, 0)), np(np(np(4, 5), np(2, 6)), np(9, 5))),
	np(7, np(np(np(3, 7), np(4, 3)), np(np(6, 3), np(8, 8)))),
	np(np(2, np(np(0, 8), np(3, 4))), np(np(np(6, 7), 1), np(7, np(1, 6)))),
	np(np(np(np(2, 4), 7), np(6, np(0, 5))), np(np(np(6, 8), np(2, 8)), np(np(2, 1), np(4, 5)))),
	np(7, np(5, np(np(3, 8), np(1, 4)))),
	np(np(2, np(2, 2)), np(8, np(8, 1))),
	np(2, 9),
	np(1, np(np(np(9, 3), 9), np(np(9, 0), np(0, 7)))),
	np(np(np(5, np(7, 4)), 7), 1),
	np(np(np(np(4, 2), 2), 6), np(8, 7)),
}

var example = []*node{
	np(np(np(0, np(5, 8)), np(np(1, 7), np(9, 6))), np(np(4, np(1, 2)), np(np(1, 4), 2))),
	np(np(np(5, np(2, 8)), 4), np(5, np(np(9, 9), 0))),
	np(6, np(np(np(6, 2), np(5, 6)), np(np(7, 6), np(4, 7)))),
	np(np(np(6, np(0, 7)), np(0, 9)), np(4, np(9, np(9, 0)))),
	np(np(np(7, np(6, 4)), np(3, np(1, 3))), np(np(np(5, 5), 1), 9)),
	np(np(6, np(np(7, 3), np(3, 2))), np(np(np(3, 8), np(5, 7)), 4)),
	np(np(np(np(5, 4), np(7, 7)), 8), np(np(8, 3), 8)),
	np(np(9, 3), np(np(9, 9), np(6, np(4, 9)))),
	np(np(2, np(np(7, 7), 7)), np(np(5, 8), np(np(9, 3), np(0, 2)))),
	np(np(np(np(5, 2), 5), np(8, np(3, 7))), np(np(5, np(7, 5)), np(4, 4))),
}

var data = []*node{
	np(np(np(0, 6), np(np(8, 9), np(3, 7))), np(np(np(3, 4), np(7, 0)), np(np(6, 9), np(4, 8)))),
	np(np(2, 2), np(np(np(7, 7), 5), np(np(0, 7), 2))),
	np(6, np(9, np(np(7, 9), 7))),
	np(np(np(np(5, 1), np(9, 3)), 8), np(4, np(2, np(6, 6)))),
	np(np(np(4, 3), np(0, 4)), np(np(np(4, 5), np(9, 3)), 3)),
	np(np(np(np(2, 7), 7), np(np(6, 5), 6)), np(np(np(2, 3), np(7, 9)), np(0, 3))),
	np(np(np(3, np(6, 2)), np(7, np(9, 4))), 3),
	np(np(np(np(9, 3), 4), np(3, 9)), 8),
	np(np(np(7, 8), np(np(2, 6), 1)), np(np(np(1, 7), 5), np(np(5, 6), np(6, 1)))),
	np(np(np(np(0, 7), 9), np(np(6, 6), np(8, 4))), np(np(np(9, 2), np(4, 8)), np(np(8, 5), np(0, 6)))),
	np(np(6, np(np(5, 6), np(3, 8))), np(np(8, 9), np(4, 3))),
	np(np(np(np(0, 6), 1), np(np(2, 4), np(1, 4))), np(np(7, 5), np(8, 3))),
	np(np(np(np(0, 7), 1), np(np(5, 7), 7)), np(np(np(3, 3), np(6, 7)), np(np(2, 8), np(2, 9)))),
	np(np(7, 7), np(np(1, np(3, 7)), 9)),
	np(np(8, np(np(3, 0), 0)), np(np(np(8, 3), 0), 9)),
	np(np(np(np(6, 2), np(2, 6)), 3), np(6, np(np(4, 7), 2))),
	np(np(np(5, np(2, 3)), np(8, np(8, 7))), np(np(0, 0), 2)),
	np(np(1, 6), np(7, np(7, np(9, 0)))),
	np(np(np(7, np(7, 6)), np(7, 4)), np(np(7, 2), np(6, 5))),
	np(1, np(np(8, np(9, 5)), 2)),
	np(np(np(np(8, 2), np(6, 5)), np(4, np(9, 2))), np(np(0, np(2, 6)), np(6, 6))),
	np(np(1, np(np(7, 2), 5)), np(np(np(6, 0), np(8, 1)), 8)),
	np(np(np(np(0, 6), np(6, 6)), 2), np(np(4, 2), np(2, 4))),
	np(np(5, np(9, 0)), np(2, 5)),
	np(7, np(np(9, 7), np(np(9, 9), 4))),
	np(np(5, np(np(6, 4), 7)), np(8, np(np(4, 4), np(9, 0)))),
	np(2, np(np(np(3, 2), np(1, 9)), np(np(3, 8), np(7, 5)))),
	np(np(np(np(8, 2), 0), np(5, np(4, 3))), 0),
	np(np(np(0, np(7, 8)), np(np(9, 6), 7)), np(np(7, np(1, 0)), np(np(0, 3), 7))),
	np(np(np(np(8, 3), 0), np(np(4, 8), np(7, 9))), np(np(7, 1), np(np(8, 4), np(4, 4)))),
	np(np(np(2, 0), np(np(6, 6), 7)), np(np(2, np(3, 9)), np(np(5, 6), np(4, 6)))),
	np(np(np(np(1, 4), 8), np(9, 6)), 8),
	np(np(7, np(9, 1)), np(1, np(np(8, 5), np(6, 8)))),
	np(8, np(np(2, 6), 5)),
	np(np(np(9, np(7, 8)), np(np(7, 8), 6)), 3),
	np(1, np(np(np(2, 1), 7), np(np(2, 6), 7))),
	np(np(7, np(4, np(6, 1))), np(np(np(4, 9), 8), np(np(0, 1), np(1, 7)))),
	np(np(np(7, 9), np(np(2, 6), np(2, 4))), np(np(2, np(1, 7)), np(np(3, 9), np(8, 9)))),
	np(np(np(np(4, 5), np(4, 7)), np(np(4, 0), np(9, 9))), 0),
	np(3, np(np(np(6, 9), 2), np(5, 3))),
	np(1, np(8, np(np(0, 8), np(1, 3)))),
	np(np(np(7, np(9, 2)), np(4, np(0, 3))), 2),
	np(3, np(np(np(7, 7), 6), np(np(8, 4), 1))),
	np(np(np(np(6, 3), np(2, 6)), np(np(6, 9), np(8, 1))), np(np(np(2, 1), np(7, 5)), np(np(7, 3), np(7, 3)))),
	np(np(np(1, 6), np(np(5, 1), np(5, 0))), np(np(1, 0), np(6, 9))),
	np(np(np(np(8, 6), np(3, 3)), np(np(2, 1), np(4, 1))), np(1, np(np(7, 7), np(8, 5)))),
	np(np(1, 5), np(6, np(np(2, 3), np(2, 4)))),
	np(np(0, np(7, np(9, 0))), np(9, 0)),
	np(np(np(5, np(1, 9)), np(0, np(9, 8))), np(np(np(6, 7), np(6, 3)), np(8, 1))),
	np(np(np(4, 7), np(6, np(2, 1))), 5),
	np(np(3, np(4, 0)), np(2, np(4, 5))),
	np(np(np(4, 0), np(6, np(8, 3))), np(np(0, 6), 8)),
	np(np(np(np(9, 9), 0), np(np(1, 8), 0)), np(np(1, 6), np(3, 4))),
	np(np(np(np(4, 3), 4), 1), np(0, np(np(2, 1), np(3, 9)))),
	np(np(np(8, np(6, 2)), np(6, 0)), 7),
	np(np(9, np(6, np(3, 1))), np(np(np(5, 9), 0), np(4, 5))),
	np(4, np(7, np(np(2, 5), 4))),
	np(np(2, np(8, np(2, 9))), np(np(np(0, 1), np(3, 5)), 1)),
	np(np(np(7, 9), np(7, 3)), np(np(1, np(7, 1)), np(1, 2))),
	np(np(np(7, 0), np(np(1, 0), 8)), np(np(9, np(7, 6)), np(9, np(7, 2)))),
	np(np(np(8, 1), np(np(0, 6), 2)), np(9, np(np(1, 8), np(5, 4)))),
	np(6, np(np(np(9, 5), np(5, 4)), 3)),
	np(np(4, np(np(6, 8), np(8, 3))), np(np(9, np(0, 9)), 7)),
	np(np(np(6, 9), np(np(2, 3), 8)), np(np(9, np(5, 1)), np(np(7, 6), 5))),
	np(np(0, 1), 5),
	np(np(4, np(1, 9)), np(np(8, 0), 8)),
	np(np(5, np(0, 6)), np(1, 8)),
	np(np(np(np(9, 2), 7), 7), np(4, np(1, np(5, 6)))),
	np(np(7, np(9, np(6, 5))), np(np(6, 9), 1)),
	np(np(np(5, 2), np(0, np(1, 4))), np(np(0, 4), np(np(9, 4), 8))),
	np(np(np(np(7, 1), np(4, 9)), 3), np(np(np(4, 5), 8), np(7, np(0, 4)))),
	np(np(np(9, np(8, 0)), 7), np(np(np(4, 5), 8), np(np(4, 3), np(8, 5)))),
	np(np(9, np(7, 0)), np(np(3, np(1, 7)), np(np(7, 0), 7))),
	np(np(2, np(np(6, 2), 6)), 8),
	np(np(np(8, np(9, 6)), np(np(5, 8), np(7, 2))), np(4, np(9, 9))),
	np(np(np(np(0, 5), 0), np(np(8, 4), 4)), np(np(7, 9), 8)),
	np(np(np(0, np(0, 3)), np(0, np(8, 8))), np(np(np(2, 1), 3), 4)),
	np(0, np(np(4, 1), np(np(9, 9), 2))),
	np(np(3, np(7, np(6, 7))), np(0, 2)),
	np(7, 2),
	np(0, np(3, np(np(3, 4), np(4, 4)))),
	np(np(np(np(0, 1), np(5, 9)), np(np(4, 2), 7)), np(5, np(1, 8))),
	np(np(7, 1), np(np(1, np(9, 9)), np(np(8, 4), 8))),
	np(np(np(1, np(8, 3)), np(np(3, 7), 0)), np(np(2, 0), np(np(1, 6), np(9, 9)))),
	np(np(np(1, 4), np(1, 4)), np(np(2, np(2, 7)), np(2, np(7, 1)))),
	np(np(1, np(np(6, 8), np(8, 6))), np(0, np(8, 0))),
	np(1, np(np(2, 0), 7)),
	np(np(np(np(6, 0), 9), np(np(6, 9), np(8, 3))), np(np(3, np(9, 9)), 6)),
	np(np(np(np(9, 8), np(2, 8)), np(2, 3)), np(6, 2)),
	np(np(np(6, np(2, 2)), 7), np(np(3, np(7, 8)), 7)),
	np(np(np(5, np(3, 7)), 1), np(np(np(4, 0), 3), np(5, 4))),
	np(np(np(7, np(4, 3)), np(9, np(4, 4))), 7),
	np(np(2, np(np(1, 5), 6)), np(np(2, 3), np(np(2, 5), np(7, 1)))),
	np(np(np(np(3, 9), np(1, 9)), 3), np(5, np(np(0, 6), np(3, 2)))),
	np(np(np(3, np(7, 5)), np(np(7, 7), np(2, 8))), np(4, np(1, np(0, 0)))),
	np(np(4, np(2, np(8, 7))), np(np(np(0, 5), 0), 9)),
	np(9, np(9, np(6, 4))),
	np(np(5, np(np(4, 9), 2)), np(9, 9)),
	np(np(1, np(np(6, 0), np(9, 9))), np(np(np(8, 4), 1), np(np(5, 2), np(6, 1)))),
	np(np(1, np(np(9, 0), 8)), 6),
}
