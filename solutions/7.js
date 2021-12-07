export function formatInput(rawInput) {
	return rawInput.trim().split(",").map(Number);
}

export function solution1(input) {
	const median = input.sort((a, b) => a - b)[Math.round(input.length / 2)];
	return input.reduce((s, n) => s + Math.abs(median - n), 0);
}

export function solution2(input) {
	let min = Infinity;
	for (let i = 0; i <= Math.max(...input); i++) {
		let sum = input.reduce((s, n) => s + (Math.abs(i - n) * (Math.abs(i - n) + 1)) / 2, 0);
		min = Math.min(min, sum);
	}
	return min;
}

export const input = `
16,1,2,0,4,2,7,1,2,14
`;

export const output = {
	part1: 37,
	part2: 168,
};
