import type { IContest } from '$lib/interfaces/contest';

function countResults(jsonArray: IContest[]): {} {
	// Initialize counters for different result types
	let wins = 0;
	let draws = 0;
	let losses = 0;

	// Iterate over each object in the JSON array
	jsonArray.forEach((contest) => {
		// Iterate over each outcome in the contest
		contest.outcomes.forEach((outcome) => {
			// Check the result and update the counters
			switch (outcome.result) {
				case 'won':
					wins++;
					break;
				case 'draw':
					draws++;
					break;
				case 'lost':
					losses++;
					break;
				case 'tie':
					draws++;
					break;
				// Add more cases if you have other result types
			}
		});
	});

	// Return an object with the result counts
	return [
		{ status: 'wins', value: wins },
		{ status: 'losses', value: losses },
		{ status: 'draws', value: draws }
	];
}

export { countResults };
