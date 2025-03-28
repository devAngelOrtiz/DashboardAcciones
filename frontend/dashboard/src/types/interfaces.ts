export interface Stock {
	ticker: string;
	target_from: string;
	target_to: string;
	company: string;
	action: string;
	brokerage: string;
	rating_from: string;
	rating_to: string;
	time: string;
}

export interface PaginatedStocks {
	stocks: Stock[];
	total: number;
	pages: number;
	page: number;
}

export interface Paginate {
	currentPage: number;
	currentSearch: string;
}
