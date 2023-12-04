import { IAccountProps, IAccount } from '../../types/AccountTypes';
import { Card, Avatar, Text, Group } from '@mantine/core';
import classes from './accountsComponent.module.css';

export function SalesDept(props: IAccountProps) {
	return (
		<div>
			<h1>Sales Department</h1>
			<div className={classes.cardContainer}>
				{props.accountsArray.map((account: IAccount) => (
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
						<Text ta="center" fz="lg" m="md">
							{account.department_name}
						</Text>
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
				))}{' '}
			</div>
		</div>
	);
}
