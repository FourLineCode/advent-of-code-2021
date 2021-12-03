export function formatInput(rawInput) {
	return rawInput.split("\n").filter(Boolean);
}

export function solution1(input) {
	let gamma = "";
	let epsilon = "";
	for (let i = 0; i < input[0].length; i++) {
		let count = [0, 0];
		for (let j = 0; j < input.length; j++) {
			if (input[j][i] === "0") count[0]++;
			else count[1]++;
		}
		gamma += count[0] > count[1] ? "0" : "1";
		epsilon += count[0] > count[1] ? "1" : "0";
	}
	return parseInt(gamma, 2) * parseInt(epsilon, 2);
}

export function solution2(input) {
	let oxygenIn = [...input];
	let co2In = [...input];
	let index = 0;
	while (oxygenIn.length > 1 || co2In.length > 1) {
		if (oxygenIn.length > 1) {
			const oxCount = [0, 0];
			for (let i = 0; i < oxygenIn.length; i++) {
				if (oxygenIn[i][index] === "0") oxCount[0]++;
				else oxCount[1]++;
			}
			if (oxCount[0] > oxCount[1]) oxygenIn = oxygenIn.filter((v) => v[index] !== "1");
			else oxygenIn = oxygenIn.filter((v) => v[index] !== "0");
		}

		if (co2In.length > 1) {
			const co2Count = [0, 0];
			for (let i = 0; i < co2In.length; i++) {
				if (co2In[i][index] === "0") co2Count[0]++;
				else co2Count[1]++;
			}
			if (co2Count[0] > co2Count[1]) co2In = co2In.filter((v) => v[index] !== "0");
			else co2In = co2In.filter((v) => v[index] !== "1");
		}

		index++;
	}
	return parseInt(oxygenIn[0], 2) * parseInt(co2In[0], 2);
}

const input = `
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`;

export const test = {
	part1: {
		input,
		output: 198,
	},
	part2: {
		input,
		output: 230,
	},
};
