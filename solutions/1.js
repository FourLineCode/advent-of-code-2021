export function formatInput(rawInput) {
	return rawInput
		.trim()
		.split("\n")
		.map((v) => +v);
}

export function solution1(input) {
	return input.reduce((a, _, i) => a + (input[i] > input[i - 1] && i > 0 ? 1 : 0), 0);
}

export function solution2(input) {
	return input.reduce((a, _, i) => {
		if (i < 3) return a;
		const prev = input.slice(i - 3, i).reduce((s, n) => s + n, 0);
		const curr = input.slice(i - 2, i + 1).reduce((s, n) => s + n, 0);
		return curr > prev ? a + 1 : a;
	}, 0);
}

const input = `
199
200
208
210
200
207
240
269
260
263
`;

export const test = {
	part1: {
		input,
		output: 7,
	},
	part2: {
		input,
		output: 5,
	},
};
