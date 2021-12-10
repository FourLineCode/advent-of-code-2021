export function formatInput(rawInput) {
	return rawInput.trim().split("\n");
}

export function solution1(input) {
	let total = 0;
	const points = { ")": 3, "]": 57, "}": 1197, ">": 25137 };
	const opening = { ")": "(", "]": "[", "}": "{", ">": "<" };
	for (const line of input) {
		const stack = [];
		for (const c of line) {
			if ("({[<".includes(c)) stack.push(c);
			else if (opening[c] === stack.pop()) continue;
			else {
				total += points[c];
				break;
			}
		}
	}
	return total;
}

export function solution2(input) {
	const scores = [];
	const points = { ")": 1, "]": 2, "}": 3, ">": 4 };
	const opening = { ")": "(", "]": "[", "}": "{", ">": "<" };
	const closing = { "(": ")", "[": "]", "{": "}", "<": ">" };
	for (const line of input) {
		const stack = [];
		let incomplete = true;
		for (const c of line) {
			if ("({[<".includes(c)) stack.push(c);
			else if (opening[c] === stack.pop()) continue;
			else {
				incomplete = false;
				break;
			}
		}
		if (!incomplete) continue;
		let score = 0;
		while (stack.length) score = score * 5 + points[closing[stack.pop()]];
		scores.push(score);
	}
	return scores.sort((a, b) => a - b)[Math.floor(scores.length / 2)];
}

export const input = `
[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
`;

export const output = {
	part1: 26397,
	part2: 288957,
};
