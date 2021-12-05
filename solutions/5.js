export function formatInput(rawInput) {
	return rawInput
		.trim()
		.split("\n")
		.map((line) => line.split(" -> ").map((c) => c.split(",").map((n) => +n)));
}

export function solution1(input) {
	const grid = [];
	for (const _ of Array(1000)) grid.push(Array(1000).fill(0));
	for (const [[x1, y1], [x2, y2]] of input) {
		if (x1 === x2) {
			const [min, max] = y2 > y1 ? [y1, y2] : [y2, y1];
			for (let i = min; i <= max; i++) grid[x1][i]++;
		} else if (y1 === y2) {
			const [min, max] = x2 > x1 ? [x1, x2] : [x2, x1];
			for (let i = min; i <= max; i++) grid[i][y1]++;
		}
	}
	return grid.reduce((a, m) => a + m.reduce((b, n) => (n > 1 ? b + 1 : b), 0), 0);
}

export function solution2(input) {
	const grid = [];
	for (const _ of Array(1000)) grid.push(Array(1000).fill(0));
	for (let [[x1, y1], [x2, y2]] of input) {
		if (x1 === x2) {
			const [min, max] = y2 > y1 ? [y1, y2] : [y2, y1];
			for (let i = min; i <= max; i++) grid[x1][i]++;
		} else if (y1 === y2) {
			const [min, max] = x2 > x1 ? [x1, x2] : [x2, x1];
			for (let i = min; i <= max; i++) grid[i][y1]++;
		} else {
			while (x1 !== x2 && y1 !== y2) {
				grid[x1][y1]++;
				x1 += x1 < x2 ? 1 : -1;
				y1 += y1 < y2 ? 1 : -1;
			}
			grid[x1][y1]++;
		}
	}
	return grid.reduce((a, m) => a + m.reduce((b, n) => (n > 1 ? b + 1 : b), 0), 0);
}

const input = `
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`;

export const test = {
	part1: {
		input,
		output: 5,
	},
	part2: {
		input,
		output: 12,
	},
};
