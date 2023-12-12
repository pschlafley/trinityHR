import { TextInput, Button, Group, Box } from '@mantine/core';
import styles from './pages.module.css';
import { useForm } from '@mantine/form';
import { ENDPOINT } from './accounts';
import { setToken } from '../utils/Authentication/Auth';

export const LoginPage = () => {
	const form = useForm({
		initialValues: {
			Email: '',
			Password: '',
		},
		validate: {
			Email: (value: string) =>
				/^\S+@\S+$/.test(value) ? null : 'Invalid email',
			Password: (value: string) =>
				/^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$/.test(
					value
				)
					? null
					: 'Invalid Password',
		},
	});

	const submitForm = async () => {
		const login = await fetch(`${ENDPOINT}/api/login`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(form.values),
		}).then((r) => r.json());

		setToken(login.token);

		form.reset();

		window.location.replace(`http://localhost:5173/accounts`);
	};

	return (
		<Box maw={340} className={styles.container} mx="auto">
			<h1>Login</h1>
			<form onSubmit={form.onSubmit(submitForm)}>
				<TextInput
					withAsterisk
					label="Email"
					placeholder="your@email.com"
					{...form.getInputProps('Email')}
				/>

				<TextInput
					withAsterisk
					label="Password"
					placeholder="Your Password"
					{...form.getInputProps('Password')}
				/>

				<Group justify="flex-end" mt="md">
					<Button type="submit">Submit</Button>
				</Group>
			</form>
		</Box>
	);
};
