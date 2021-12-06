export function formatInput(rawInput) {
	const input = rawInput.trim().split("\n\n");
	const numbers = input[0].split(",").map((n) => +n);
	const boards = input.slice(1).map((grid) =>
		grid.split("\n").map((v) =>
			v
				.split(" ")
				.filter(Boolean)
				.map((n) => ({ value: +n, marked: false }))
		)
	);

	return { numbers, boards };
}

export function solution1({ numbers, boards }) {
	const sum = (board) => {
		let s = 0;
		for (const c of board) for (const x of c) s += !x.marked ? x.value : 0;
		return s;
	};

	for (const n of numbers) {
		for (const board of boards) {
			for (const row of board) {
				for (const col of row) {
					if (col.value === n) {
						col.marked = true;
					}
				}
			}
			let won = false;
			for (let i = 0; i < board.length; i++) {
				let found = true;
				for (let j = 0; j < board[i].length; j++) {
					if (!board[i][j].marked) {
						found = false;
						break;
					}
				}
				if (found) {
					won = true;
					break;
				}
			}
			for (let i = 0; i < board.length; i++) {
				let found = true;
				for (let j = 0; j < board[i].length; j++) {
					if (!board[j][i].marked) {
						found = false;
						break;
					}
				}
				if (found) {
					won = true;
					break;
				}
			}
			if (won) return sum(board) * n;
		}
	}
}

export function solution2({ numbers, boards }) {
	const sum = (board) => {
		let s = 0;
		for (const c of board) for (const x of c) s += !x.marked ? x.value : 0;
		return s;
	};

	const winning = [];
	for (const n of numbers) {
		for (const [k, board] of boards.entries()) {
			if (winning.includes(k)) continue;
			for (const row of board) {
				for (const col of row) {
					if (col.value === n && !col.marked) {
						col.marked = true;
					}
				}
			}
			let won = false;
			for (let i = 0; i < board.length; i++) {
				let found = true;
				for (let j = 0; j < board[i].length; j++) {
					if (!board[i][j].marked) {
						found = false;
						break;
					}
				}
				if (found) won = true;
			}
			for (let i = 0; i < board.length; i++) {
				let found = true;
				for (let j = 0; j < board[i].length; j++) {
					if (!board[j][i].marked) {
						found = false;
						break;
					}
				}
				if (found) won = true;
			}
			if (won) {
				if (boards.length - winning.length > 1) winning.push(k);
				else return sum(board) * n;
			}
		}
	}
}

export const input = `
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`;

export const output = {
	part1: 4512,
	part2: 1924,
};
