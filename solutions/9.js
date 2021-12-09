export function formatInput(rawInput) {
	return rawInput
		.trim()
		.split("\n")
		.map((l) => l.split("").map(Number));
}

export function solution1(input) {
	let risk = 0;
	const inBound = (x, y) => x >= 0 && x < input.length && y >= 0 && y < input[x].length;
	for (let i = 0; i < input.length; i++) {
		for (let j = 0; j < input[i].length; j++) {
			let min = Infinity;
			const adjacent = [
				[i + 1, j],
				[i - 1, j],
				[i, j + 1],
				[i, j - 1],
			];
			for (const [x, y] of adjacent) if (inBound(x, y)) min = Math.min(min, input[x][y]);
			if (input[i][j] < min) risk += input[i][j] + 1;
		}
	}
	return risk;
}

export function solution2(input) {
	const basins = [];
	const inBound = (x, y) => x >= 0 && x < input.length && y >= 0 && y < input[x].length;
	for (let i = 0; i < input.length; i++) {
		for (let j = 0; j < input[i].length; j++) {
			let min = Infinity;
			const adjacent = [
				[i + 1, j],
				[i - 1, j],
				[i, j + 1],
				[i, j - 1],
			];
			for (const [x, y] of adjacent) if (inBound(x, y)) min = Math.min(min, input[x][y]);
			if (input[i][j] < min) {
				let size = 0;
				const visited = [];
				const stack = [[i, j]];
				while (stack.length) {
					const [x, y] = stack.pop();
					const sides = [
						[x + 1, y],
						[x - 1, y],
						[x, y + 1],
						[x, y - 1],
					];
					for (const [m, n] of sides) {
						if (inBound(m, n) && input[m][n] !== 9) {
							if (!visited.includes(String([m, n]))) {
								visited.push(String([m, n]));
								stack.push([m, n]);
								size++;
							}
						}
					}
				}
				basins.push(size);
			}
		}
	}
	return basins.sort((a, b) => b - a).reduce((a, n, i) => (i < 3 ? a * n : a), 1);
}

export const input = `
2199943210
3987894921
9856789892
8767896789
9899965678
`;

export const output = {
	part1: 15,
	part2: 1134,
};
