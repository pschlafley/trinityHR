/* eslint-disable no-mixed-spaces-and-tabs */
import classes from './accounts.module.css';
import { Card, Avatar, Text, Group } from '@mantine/core';
import useSWR from 'swr';
import Account from '../../types/AccountTypes';

export const ENDPOINT = 'http://localhost:3000';

const fetcher = async (url: string) => {
	const res = await fetch(`${ENDPOINT}/${url}`);
	return await res.json();
};

export const AccountsComponent = () => {
	const arrayData: Account[] = [];

	const { data } = useSWR('api/accounts', fetcher);

	for (let i = 0; i < data?.length; i++) {
		arrayData.push(data[i]);
	}
	return (
		<div className={classes.container}>
			{arrayData.map((account: Account) => (
				<Card
					key={account.account_id}
					withBorder
					padding="xl"
					radius="md"
					className={classes.card}
				>
					<Card.Section
						h={140}
						style={{
							backgroundImage:
								'url(https://images.unsplash.com/photo-1488590528505-98d2b5aba04b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=500&q=80)',
						}}
					/>
					<Avatar
						src="https://raw.githubusercontent.com/mantinedev/mantine/master/.demo/avatars/avatar-9.png"
						size={80}
						radius={80}
						mx="auto"
						mt={-30}
						className={classes.avatar}
					/>
					<Text ta="center" fz="lg" fw={500} mt="sm">
						{account.full_name}
					</Text>
					<Text ta="center" fz="sm" c="dimmed">
						{account.role}
					</Text>
					<Group mt="md" justify="center" gap={30}>
						{account.account_type}
					</Group>
				</Card>
			))}
		</div>
	);
};
