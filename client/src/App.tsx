/* eslint-disable no-mixed-spaces-and-tabs */
import '@mantine/core/styles.css';
import useSWR from 'swr';
import Account from './types/AccountTypes';
import { HeroComponent } from './components/heroComponent/HeroComponent';
import { AccountsComponent } from './components/Accounts/accounts';
import { NavBar } from './components/Navbar/navbar';
import { MantineProvider } from '@mantine/core';

export const ENDPOINT = 'http://localhost:3000';

const fetcher = async (url: string) => {
	const res = await fetch(`${ENDPOINT}/${url}`);
	return await res.json();
};

function App() {
	const arrayData: Account[] = [];

	const { data } = useSWR('api/accounts', fetcher);

	for (let i = 0; i < data?.length; i++) {
		arrayData.push(data[i]);
	}

	return (
		<div>
			<MantineProvider>
				<NavBar />
				<HeroComponent />
				<AccountsComponent data={data} arrayData={arrayData} />
			</MantineProvider>
		</div>
	);
}

export default App;
