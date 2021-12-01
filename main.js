import fs from "fs";

console.log("\n");

const [, , day] = process.argv;
if (!day) {
	throw new Error("Please provide the day of the solution!");
}

const { formatInput, solution1, solution2 } = await import(`./solutions/${day}.js`);
const dayInput = await fs.promises.readFile(`./inputs/${day}.txt`, "utf8");

if (!solution1) {
	console.warn('No solution 1 found. Make sure your module exports a "solution1" function.');
} else {
	console.log("Running solution 1 ...");
	console.time("Time:");
	console.log("Solution - 1:", await solution1(await formatInput(dayInput)));
	console.timeEnd("Time:");
	console.log();
}

if (!solution2) {
	console.warn('No solution 2 found. Make sure your module exports a "solution2" function.');
} else {
	console.log("Running solution 2 ...");
	console.time("Time:");
	console.log("Solution - 2:", await solution2(await formatInput(dayInput)));
	console.timeEnd("Time:");
	console.log();
}
