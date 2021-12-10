import fs from "fs";

const [, , day] = process.argv;
if (!day) throw new Error("Please provide the day of the solution!");

const { formatInput, solution1, solution2, input, output } = await import(`./solutions/${day}.js`);
const dayInput = await fs.promises.readFile(`./inputs/${day}.txt`, "utf8");

if (!solution1) {
	console.warn("No solution 1 found!");
} else {
	console.log("Running solution 1 ...");
	const result = await solution1(await formatInput(input));
	const passed = result === output.part1;
	console.log("Test for part 1 -", passed ? "\x1b[32mPASSED\x1b[0m" : "\x1b[31mFAILED\x1b[0m");
	if (!passed) console.log("Expected:", output.part1, "| Got:", result);
	console.log("Solution - 1:", await solution1(await formatInput(dayInput)));
	console.log();
}

if (!solution2) {
	console.warn("No solution 2 found!");
} else {
	console.log("Running solution 2 ...");
	const result = await solution2(await formatInput(input));
	const passed = result === output.part2;
	console.log("Test for part 2 -", passed ? "\x1b[32mPASSED\x1b[0m" : "\x1b[31mFAILED\x1b[0m");
	if (!passed) console.log("Expected:", output.part2, "| Got:", result);
	console.log("Solution - 2:", await solution2(await formatInput(dayInput)));
}
