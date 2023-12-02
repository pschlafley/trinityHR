export interface IAccount {
	account_id: string;
	account_type: string;
	role: string;
	full_name: string;
	email: string;
	password: string;
	department_id: number;
	department_name: string;
}

export interface IAccountProps {
	accountsArray: IAccount[];
}
