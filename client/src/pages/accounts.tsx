/* eslint-disable no-mixed-spaces-and-tabs */
import classes from './accounts.module.css';
import pageStyles from './pages.module.css';
import useSWR from 'swr';
import { IAccount } from '../types/AccountTypes';
import { AssemblyInstallationDept } from '../components/Accounts/Installation';
import { SalesDept } from '../components/Accounts/Sales';
import { ManagementDept } from '../components/Accounts/Management';
import Cookies from 'js-cookie';

export const ENDPOINT = 'http://localhost:3000';

const user = Cookies.get('token');

const fetcher = async (url: string) => {
	const res = await fetch(`${ENDPOINT}/${url}`, {
		headers: {
			'x-jwt-token': user ? user.toString() : '',
		},
	});
	return await res.json();
};

export const AccountsComponent = () => {
	const assemblyInstallArray: IAccount[] = [];
	const salesArray: IAccount[] = [];
	const managementArray: IAccount[] = [];

	const accountReq = useSWR('api/accounts', fetcher, {
		refreshInterval: 60000,
	});
	const departmentReq = useSWR('api/departments', fetcher, {
		refreshInterval: 60000,
	});

	for (let i = 0; i < accountReq.data?.length; i++) {
		for (let j = 0; j < departmentReq.data?.length; j++) {
			if (
				departmentReq.data[j].department_name == 'Assembly and Installation' &&
				accountReq.data[i].department_id == departmentReq.data[j].department_id
			) {
				assemblyInstallArray.push(accountReq.data[i]);
			} else if (
				departmentReq.data[j].department_name == 'Sales' &&
				accountReq.data[i].department_id == departmentReq.data[j].department_id
			) {
				salesArray.push(accountReq.data[i]);
			} else if (
				departmentReq.data[j].department_name == 'Management' &&
				accountReq.data[i].department_id == departmentReq.data[j].department_id
			) {
				managementArray.push(accountReq.data[i]);
			}
		}
	}

	return (
		<div className={`${classes.container} ${pageStyles.container}`}>
			{user ? (
				<div>
					<AssemblyInstallationDept accountsArray={assemblyInstallArray} />
					<SalesDept accountsArray={salesArray} />
					<ManagementDept accountsArray={managementArray} />
				</div>
			) : (
				<div>Not Authenticated</div>
			)}
		</div>
	);
};
