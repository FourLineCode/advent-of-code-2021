import dotenv from 'dotenv';
import fs from 'fs';
import fetch from 'node-fetch';

dotenv.config();

const [, , day] = process.argv;
if (!day) {
	throw new Error('Please provide the day of the input!');
}

const cookie = process.env.AOC_COOKIE;
if (!cookie) {
	throw new Error('Please provide AOC session cookie in environment variable!');
}

const response = await fetch(`https://adventofcode.com/2021/day/${day}/input`, {
	headers: { cookie },
});
const contents = await response.text();

await fs.promises.writeFile(`./inputs/${day}.txt`, contents, 'utf8');
console.log(`Successfully fetched input for day - ${day}`);
