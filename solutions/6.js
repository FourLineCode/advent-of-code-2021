export function formatInput(rawInput) {
	return rawInput.trim().split(",").map(Number);
}

export function solution1(input) {
	let fish = {};
	for (const n of input) fish[n] = n in fish ? fish[n] + 1 : 1;
	for (const _ of Array(80)) {
		const tmp = {};
		for (let [day, count] of Object.entries(fish)) {
			day = parseInt(day);
			if (day === 0) {
				tmp[6] = 6 in tmp ? tmp[6] + count : count;
				tmp[8] = 8 in tmp ? tmp[8] + count : count;
			} else {
				tmp[day - 1] = day - 1 in tmp ? tmp[day - 1] + count : count;
			}
		}
		fish = { ...tmp };
	}
	return Object.values(fish).reduce((a, n) => a + n, 0);
}

export function solution2(input) {
	let fish = {};
	for (const n of input) fish[n] = n in fish ? fish[n] + 1 : 1;
	for (const _ of Array(256)) {
		const tmp = {};
		for (let [day, count] of Object.entries(fish)) {
			day = parseInt(day);
			if (day === 0) {
				tmp[6] = 6 in tmp ? tmp[6] + count : count;
				tmp[8] = 8 in tmp ? tmp[8] + count : count;
			} else {
				tmp[day - 1] = day - 1 in tmp ? tmp[day - 1] + count : count;
			}
		}
		fish = { ...tmp };
	}
	return Object.values(fish).reduce((a, n) => a + n, 0);
}

export const input = `
3,4,3,1,2
`;

export const output = {
	part1: 5934,
	part2: 26984457539,
};
