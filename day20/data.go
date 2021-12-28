package main

var testAlgo = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#`

var testImage = []string{
	"#..#.",
	"#....",
	"##..#",
	"..#..",
	"..###",
}

var algo = `#####..#.#######.##.#.#.#.#####..######...#.##.###...#.##..#####..###.#.####..#.##.#....#...##.###..#.#......####.#...#.####.#..#.###.#.#.#.###.###..########.##.#.#....#.#####.......###.#..#.#.###.###.###.#.......#.#....#.##.###.....##...#.#.#..#.#....##..#####...##...#.##..##.##.#...###...#.##...#...#.####.###...........#...#..##...##.###.##.###.....#.##...###.....#.#.###..##..#..##..#.....##..##....#....####...#.###.#######.#.##.####.....####.#.#.##.#####...#.##.##.#.##..####.#.#.######.#..###..#.##.##...`

var image = []string{
	"##.##....#..#....#.....#..##.#...######.##.##..#..##.###.##..##.#..#..####..##..#.#..###..##...##.##",
	"..##.###.##..##.###..#.###....##.#.#...##.#.#..####.#..#..##...##......###...###.#..##...##.##.##..#",
	".#..####..##..###..#.....#....##....###..####.##.##.#.#.##..######..#####.#...##.#..##..##..##...##.",
	"..###.......##.#..#.##.##.#..#######..##.......###.###.###.#..###...#.#####..###....###.##..###..##.",
	".#..####..###.#.#..#.#..##...#...#.##...##..##..##.#######.....#.#..#.#.#.#.#.#..#.##......####..##.",
	"##..#...###.##....####.###.###.#.#.#........##...#.#....#######.######...##.###...##.###.##.######..",
	"##..##...#....#....#####.#.....#....#.#..#.#.##...##.##...#.####...#####.#.#.##...#..##.##.##.##..##",
	"..#..##..#..#...####....#.....###....##.#..#...#..#.##.#....#.#.#...##.#..#..#..##....#.##..#..##.##",
	"..##....######.###.#.....##.##....###...#..#.....#...###.###....#...#.##.####..##.###....#####...##.",
	".#..##.###.#......#.#####.####..####......###.....####..####..####..#.######.####..##.#...#.##.####.",
	"####..##.##.##..#.#...##.####..####.#...#.#.##..#.###.#.....#....###.####.###..#.#..#.#...#.#.##..##",
	".##..#....###..#####.##.....###.##.###.######.#.###.#.##.#.#######....##..#.#.##..#.##...#######..##",
	"###.......###.....#.#.##....##..#.##...#....####.##.###.....#.####..#........#...####.##.#.##...##.#",
	"##.##....##..#####.#...####.....##..#..#...##.#...##.#..#..########...###.##...#.##...##.#..##.##...",
	".#.###...#.###..##...####..#.###.##.#.#...#..###.#.####..#..##.##.####.###...##.##.###..#.#.####.##.",
	"#..#.#.#.#.#.#.#..#####.#....####.#.......#..#####..##..#..###....#...#..###.#..###.##..#.###.##.###",
	"....##..##.###..####.#.###.##.....#...#...##.##.#.###..##.#.#..#....#.####.##...#.#...#.#...###...#.",
	"##...#....##.#....###....#..##.##.####.#.#.##..##.####...#.#....####..#.##.....#....#...#..##...#...",
	"#.##...#.##..#.##.###..##...##.#.#..##..#.....###.#.####....#######.#.#.#####.#....##.#.....####....",
	".....#####..###...#.#.#.#...#.#..#.###..##.##....##.###..###..###.#.#...#.#.##.###...#...##.#..#.##.",
	"...###.###.###.#..#.#.####.##...##.##..#####.......#.##.....#####..#.#######...#..#.###......#######",
	"#..##.....###..#.###..#..###.#.##...#...#......#.###..#...###...#..##......#.....##.####..##...#...#",
	".....##..#.#.#..#.#.#.#####...#..#.......#..#..#....#..##..##.#.#..#.####..#...#..##...##...#..#####",
	"...##..#.#...#...#.####..#.##.###..##..###.......#..######.#.###.#..#.##.###.....########..#..#...#.",
	"#.#...##...#.##.##.#..#..#.#....##.#.#...#.#..####.#...#.#..######.#.##....#.#...##....#..#.#.##.###",
	"...#####.##.####.##...#..##.#..##...#..##..#.##...##..#.#.###......#.###.#.#..###...##..##.##..###..",
	".##...###..#....#.#.....#..#.##..###.#..###..###.#..#.#.###..#...#..####..#.###.##.##.#.#.#..#..#..#",
	".#..#.#.#..#####.##...#.#...###.####.##..###....######.##..##.#...#.###..###.###....##.#......#.##.#",
	"#..####.######....###..####.##.##.....#.##.####.##..##.#..##.#.#.#.##.##..#.###.##..####.##.#.##..##",
	"....#.#..........#.##....#.#..#.###.##.####.#.#..##.#..##.#..#.....###.#..###.######..##.#......####",
	"...#..#.##.###########...###.###..##......#..###.#.#...##.##...####.##..#.........##...##.###.###.#.",
	".##.#####.#...#.###.#..#..#.#..#.##...#.#..##.#.####.##..######....##.####...#..###.####.#...##.#...",
	"...#.##..#..#.#..##..#.#..##.....#.##.#....#..#..#.#......##.####...####...##......##..##....####.#.",
	"####.##.#...#....#.#.#..###.##.###..#.###.#.#..#.##.##..##.##.#.#..#.####...###.####.##.....#.#...##",
	"#........#.###..#.#.#####...#.##.#.###...#..#...##.##..#.###....#.###..#.#.#.###.#.##.#....#......##",
	"...#.##..##.#..........#..#....###.#..#..#.#...###.##..#.#...###.####...####..#.####.######.#....#..",
	"##.#.####.##.#.##..#..##...###...###.##..##.#####..###..#..#.#..#.##...#.####...##.##..#.#.#..###...",
	"#.####..##..#.#..##..##.#..#..#...#.###.#.##.####.###.###...##.#####..##..####...##.##########.#.###",
	"####.##..###.####.#.######...#..#.##.#.....##.#....#...#.############..#.#.###.####.#.#..#.###....#.",
	"..###..#...##.#.##.##..######.###.###..#..#####....#.#....####.#..#.......##..#####.#...#.#.##....#.",
	".#.##...#..#.##.####......#.#......######.#.#.##.#..##..#....##..##....#.##.###..#.##..#.#.....##...",
	"###..###.########..####.###.#...###.........##....##.##..#...#.##..##..#.#.####....######..#..#...##",
	"##.#..###..#.#..#...##.####.#..##..##...###.##.#......#........###...#.###..####..####..#####..###..",
	"#.##...##..##..#.#.###..###...###.....#.#.######.....##.....#.##..##.#...###.###.#.#..#.##.##.##.#..",
	"#..#.#####..#..#...#.#...#..##.#..#.###...#..###.###.#.#...#.###..#.###..#..##.##......##...#...###.",
	"#....#.##.....#.#...###....#..#.#.#..##.##..###.###..#..#.#..##.....###.#.#.#...#.##.#..##....####..",
	"#..#####...#...#####.###.#..#..#...####.##.#.#..#...#.####...##.##.....#.#.##.##.##.#....#.######..#",
	"###...##..##.##.##.#..##.##.#.##.##.#..#.#####..#.#..#.#..####....#.###..###...####.#....#...#.###.#",
	"##...####...##.######...#...#..##..###..###..#.##.#.#.#...##.#.###...#...########..#.#######...###..",
	".###...#.#.####.#.#...##.#.###.#.###.#####.###..##.#..###..#.#...##...#####..##....#..#..##.#.#....#",
	"...#..#....#..#....#.#.##..#..#####....#..#.#..####..####.####.###......###..#..#..###.##.###...###.",
	"#.#...#.###.###....#..###..###...#..#####..#.##........###....###.###.##.###.#..##..#.....#.#.####.#",
	"#...#.#...#.#..#..#.##.##.#......#.#.#..###....#.#.#.#..#.#.###..#..#..##.#.#.#.##...#..###.####..#.",
	"#.#.##.####........#.##..#.#.#...##..#.#####.#..#.#..#.##.#..####.###.#####.#.##.#....##..#####.####",
	"#######.##..#.#..#...#..#..#.######..##.#.###.#...#..####..##.#.##.###.#.##.#.##..#...#.#..#.#.####.",
	".##..#####.##.####.##.#....##.....###.......#.####.###....#...#.#####.###.#..#...##.#..##.#....###.#",
	"..####.#.#.....#...###.#...#.#...#.####..##..#..######......#...#.##.####..#.......#.##.......#.###.",
	".##..#.#####..##..##.#..####.##.##.#.###.#.#.....#.....#.#...####....##.##...####......##.##.#######",
	"#.#.#..###..###..##.#..##...###..###..#..#.###.#..##.##...##..#..####....####.##.#.###...#..##.#.##.",
	"##.##.....######.#..##.#.#.###.##.###.##....######.##.#.##..#..##...#...#..#.#.#.######..#.###.##..#",
	"#..#.#.###.#.#.#.....###.######..##.##.#.#...#####.#...##.##.##.##....##....#..#.##.#.##..##...#...#",
	"..#.....#....#...#######..#.#.#.#.###.##.#.##.###..#..##..###..#.#####..###.......####.#..####..#...",
	"#..#.#...##.##..#.#.....###.##.#.#....###....#.##..#....#####.##.#####..###...####.#.#.#.###.#....##",
	"#....##..#..##..#.####..##.##....####..##..#..####.######.##..#.###.#.##....###.#....##....##.#.##.#",
	".#.#.#.##.##...###.#.#...#...#..######.##.#.#.##.####.....#.#..#.######.......##.#####..#....#####.#",
	".###.#.##...###.....##..#.#..##....#.###..##.##..####..##.#.###..##.#.#...##.#####.#.....#..#..#####",
	"#.#.###..#....####....#...#.###.#..#...###..###......##..##.#.#..#....#......###...##..#.##.##...#.#",
	"###.####.##.#.##.......##.....#.#.###......##..#.#....###..#.#.##.###.##.##.##.#..#.####.#......###.",
	"####...#...##...#.#.#..######.#####..##...#.##.##.####.##.#.##..#..#.#.##.#...#......#.##.#######...",
	".##...#....#.##.#...#..#..#.####..#.#..........#.#.##......#.##.######.###.#####.#.#####.###.####...",
	".##.##.##.###..###..##.#..#.###..#####.#..##.##.#.####....###.#..###...##..##..#.##..#...#..###.##.#",
	"#.....##..#...##.#...#.###.##..#.....##..#.##...#.##...#.#.##.#...#.#..###...#...#######...#.#.....#",
	".######..#.####.#.##.##.#.#.#.#.###.#######...#..####.#.#..#####..#.....###.##...#.###.####.#.#.....",
	"#.####..##...####.##.#.##...####..##.#.#..####.##.#.###.##..##.##..###.##...##........###..#.#.##...",
	"#...#.##.....###.##.....####.####.##.##.##....##.#.#.#.##...#.#.##..#..##.....##.#...##..##...##..##",
	"#..#.###.#....#.##...##...#.#......##.###...#......##..#..##...##.#..###.#..#.#.#.##.........##.#.##",
	".#.#..#..#.#..##...#.#...#..###.#...#.#.##.##....#.##.#...#.##.#..#.#..##..#.##...#.#.#..###.#.###..",
	"###...#.#..##.###...#..##.#..##.#.######..####..#.#.....##....##.####.###...#.#...#.#..#..#.###....#",
	".####.###..##.##...#.....##.###.#..##..####.....#....##.....#....##..##...##....#..#..###.#...###.##",
	"..###..###.#.#....#.##........#...####...####...#.####.##.###.##...#.####.#.###..###.##...##...#.#.#",
	"###.#..#.#...#.#.#.....##.###...#..###.#....#.##.##.###..#######.###.#...###.####...#.#.###.#...####",
	"#.#....####..#.#.###..###.###.#.#.##...###.#..###.#.#.#...###..#...##...#..#..###..##.#..#.##...###.",
	".###..##.#..#.#.#####.#..#.........##..##...#...#.#...##..#.##...###..#.#.#..####....#.####.##.#....",
	"##.....#######....#.###.##...##.##.#####.#....#########.##....##.###...#.###..##.###.#..##..#....#.#",
	"...#.##.#.......#.#....#..##..#..##...#.#.####...###.##..##...##.#.##.#.###.###.###..####..#######.#",
	".#.###...##.##.#.......#.#...#..#####...##...#...#.#####.##.#......#.##...#.###.#.#..#.#..##..##.#.#",
	"...##...##.#.#...#...#.###..###....#.##...##.....#...##..##....#..##..#.#..#.#.#.........##.#####..#",
	"........#...#.......#.#..#..#......##..##...##......#.....#.#.......#.#####.#.#.#..#...#####.#..####",
	"#.#.###...###....#...######...#.#.#####.##..#..##.##..#..#..#..##...###..#.#.#..##...#.#.#.#.####..#",
	".#..#...##...#....#.#.##.#.#....##.#....#...#.#.##.....#########.#....#.##..#..#.#.##..#.#..###....#",
	"##..#.#.#..###.....###.#..#.##..##....###.#..####.#.#.#..######..###..#.#.#...#...#...#.#..#.#..###.",
	".#.##.#.####....#..#.#...#.##......#..#.....#...#....#.##########..#...#.#.###...####..###......##.#",
	"###...#....###.#.#..#...#..####.#...#.##..#.##.#...###.#.#..#.###.###...#......#..#.##.#..#####...##",
	".###.##.#..###...#..####.###...#.##..##....######.#..##...#...#.##.....#.####.########.##....#.##.##",
	".......######.#...#...#.###....##..#.###.#...##.##..###.#..##...#..########.#.#.#...#..#..##.##..#.#",
	"..###.##..#..#####.#.##.#..#...##..#.#..###.###....###...#..##......###..#....#.#.....#...#..##.#..#",
	"###..##..#.#....#####.#######..#..##...#..#...#.#........#...###.#.###..##.####..#.####.#..#..#...#.",
	"...#..###.#.#.#.###...#.......####..###.#.##.####.#..#.#..#########.#.....#...##..#.#.#..#.#.#.#...#",
	"####.#..##.##..#..###.######.#.#.###...#####.#.###.#####.#.#.#.#...#.#.#...#.#########.#..###.##.###",
	".##...#.#.###....##..##..#.#.#.#.#..#..##.###...####..#..#.#.#.#####..##.#.#.##.##...#.....#...###.#",
}
