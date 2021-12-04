import fs from "fs";

const [, , day] = process.argv;
if (!day) throw new Error("Please provide the day of the solution!");

const { formatInput, solution1, solution2, test } = await import(`./solutions/${day}.js`);
const dayInput = await fs.promises.readFile(`./inputs/${day}.txt`, "utf8");

if (!solution1) {
	console.warn("No solution 1 found!");
} else {
	console.log("Running solution 1 ...");
	const result = await solution1(await formatInput(test.part1.input));
	console.log(
		"Test for part 1 -",
		result === test.part1.output ? "\x1b[32mPASSED\x1b[0m" : "\x1b[31mFAILED\x1b[0m"
	);
	console.log("Solution - 1:", await solution1(await formatInput(dayInput)));
	console.log();
}

if (!solution2) {
	console.warn('No solution 2 found. Make sure your module exports a "solution2" function.');
} else {
	console.log("Running solution 2 ...");
	const result = await solution2(await formatInput(test.part2.input ?? test.part1.input));
	console.log(
		"Test for part 2 -",
		result === test.part2.output ? "\x1b[32mPASSED\x1b[0m" : "\x1b[31mFAILED\x1b[0m"
	);
	console.log("Solution - 2:", await solution2(await formatInput(dayInput)));
}
