export function formatInput(rawInput) {
	return rawInput
		.trim()
		.split("\n")
		.map((line) => line.split(" | ").map((p) => p.split(" ")));
}

export function solution1(input) {
	let count = 0;
	for (const digits of input.map(([_, digits]) => digits))
		for (const digit of digits) if ([2, 3, 4, 7].includes(digit.length)) count++;
	return count;
}

export function solution2(input) {
	let sum = 0;
	const pattern_list = input.map(([p]) => p.map((a) => a.split("").sort().join("")));
	const digit_list = input.map(([_, d]) => d.map((b) => b.split("").sort().join("")));
	for (let i = 0; i < input.length; i++) {
		const patterns = pattern_list[i];
		const digits = digit_list[i];
		let sixLengths = patterns.filter((p) => p.length == 6);
		let fiveLengths = patterns.filter((p) => p.length == 5);
		const found = {};

		found[1] = patterns.filter((p) => p.length == 2)[0];
		found[7] = patterns.filter((p) => p.length == 3)[0];
		found[4] = patterns.filter((p) => p.length == 4)[0];
		found[8] = patterns.filter((p) => p.length == 7)[0];

		found[6] = sixLengths.filter((p) => {
			for (const c of found[7]) if (!p.includes(c)) return true;
			return false;
		})[0];
		sixLengths = sixLengths.filter((p) => p !== found[6]);

		found[5] = fiveLengths.filter((p) => {
			let different = 0;
			for (const c of found[6]) if (!p.includes(c)) different++;
			return different === 1;
		})[0];
		fiveLengths = fiveLengths.filter((p) => p !== found[5]);

		found[2] = fiveLengths.filter((p) => {
			let similar = 0;
			for (const c of found[4]) if (!p.includes(c)) similar++;
			return similar === 2;
		})[0];
		found[3] = fiveLengths.filter((p) => p !== found[2])[0];

		found[9] = sixLengths.filter((p) => {
			for (const c of found[4]) if (!p.includes(c)) return false;
			return true;
		})[0];
		found[0] = sixLengths.filter((p) => p !== found[9])[0];

		const parsed = {};
		for (const key in found) parsed[found[key]] = key;

		sum += parseInt(digits.reduce((n, d) => n + parsed[d], ""));
	}
	return sum;
}

export const input = `
be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
`;

export const output = {
	part1: 26,
	part2: 61229,
};
