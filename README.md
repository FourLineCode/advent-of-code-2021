# Advent of Code Solutions 2021

My advent of code 2021 solutions.

## Using the setup

**Run solution for a specific day**

```bash
npm start <day>
```

**Running on dev mode for a specific day:**

```bash
npm run dev <day>
```

**Fetch input for a specific day**

```bash
npm run input <day>
```

## Defining a solution

Solutions are defined in JavaScript files in the `solutions` directory. The module should be named after the day (i.e. `<day>.js`), and should export two functions: `solution1` and `solution2`. Each of these solution functions will be called with the input file for that specific day, and the return value will be displayed to the console. The runner will automatilly watch the files and re-run whenever a code change is made. Your solution files can export async functions as well if needed.

## Solution scaffold

```javascript
function formatInput(rawInput) {
	return rawInput;
}

export function solution1(rawInput) {
	const input = formatInput(rawInput);

	return 'TODO';
}

export function solution2(rawInput) {
	const input = formatInput(rawInput);

	return 'TODO';
}
```
