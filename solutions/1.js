export function formatInput(rawInput) {
	return rawInput.split("\n").map((v) => +v);
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
