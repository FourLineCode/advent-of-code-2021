export function formatInput(rawInput) {
	return rawInput
		.trim()
		.split("\n")
		.map((l) => l.split("").map(Number));
}

export function solution1(input) {
	let flashes = 0;
	const flash = (x, y) => {
		flashes++;
		input[x][y] = -1;
		for (const m of [-1, 0, 1]) {
			for (const n of [-1, 0, 1]) {
				const [a, b] = [x + m, y + n];
				if (a < 0 || a >= input.length || b < 0 || b >= input[a].length) continue;
				if (input[a][b] === -1) continue;
				input[a][b]++;
				if (input[a][b] > 9) flash(a, b);
			}
		}
	};
	for (let d = 0; d < 100; d++) {
		for (let i = 0; i < input.length; i++)
			for (let j = 0; j < input[i].length; j++) input[i][j]++;
		for (let i = 0; i < input.length; i++)
			for (let j = 0; j < input[i].length; j++) if (input[i][j] === 10) flash(i, j);
		for (let i = 0; i < input.length; i++)
			for (let j = 0; j < input[i].length; j++) if (input[i][j] === -1) input[i][j] = 0;
	}
	return flashes;
}

export function solution2(input) {
	const flash = (x, y) => {
		input[x][y] = -1;
		for (const m of [-1, 0, 1]) {
			for (const n of [-1, 0, 1]) {
				const [a, b] = [x + m, y + n];
				if (a < 0 || a >= input.length || b < 0 || b >= input[a].length) continue;
				if (input[a][b] === -1) continue;
				input[a][b]++;
				if (input[a][b] > 9) flash(a, b);
			}
		}
	};
	for (let d = 1; ; d++) {
		for (let i = 0; i < input.length; i++)
			for (let j = 0; j < input[i].length; j++) input[i][j]++;
		for (let i = 0; i < input.length; i++)
			for (let j = 0; j < input[i].length; j++) if (input[i][j] === 10) flash(i, j);
		let flashed = true;
		for (let i = 0; i < input.length; i++)
			for (let j = 0; j < input[i].length; j++)
				if (input[i][j] === -1) input[i][j] = 0;
				else flashed = false;
		if (flashed) return d;
	}
}

export const input = `
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`;

export const output = {
	part1: 1656,
	part2: 195,
};
