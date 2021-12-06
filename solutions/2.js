export function formatInput(rawInput) {
	return rawInput
		.trim()
		.split("\n")
		.map((v) => v.trim().split(" "));
}

export function solution1(input) {
	let height = 0;
	let depth = 0;
	for (const [dir, num] of input) {
		if (dir === "forward") {
			height += +num;
		} else if (dir === "down") {
			depth += +num;
		} else {
			depth -= +num;
		}
	}
	return height * depth;
}

export function solution2(input) {
	let height = 0;
	let depth = 0;
	let aim = 0;
	for (const [dir, num] of input) {
		if (dir === "forward") {
			height += +num;
			depth += aim * +num;
		} else if (dir === "down") {
			aim += +num;
		} else {
			aim -= +num;
		}
	}
	return height * depth;
}

export const input = `
forward 5
down 5
forward 8
up 3
down 8
forward 2
`;

export const output = {
	part1: 150,
	part2: 900,
};
