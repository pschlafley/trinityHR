import { useLocation } from 'react-router-dom';
import { Container, Group, Image } from '@mantine/core';
import classes from './navbar.module.css';
import logo from '../../assets/TrinityLogo.png';

export function NavBar() {
	const windowPath = useLocation().pathname;

	const homePath = windowPath == '/' ? true : false;
	const accountsPath = windowPath == '/accounts' ? true : false;

	return (
		<header className={classes.header}>
			<Container size="md" className={classes.inner}>
				<Group gap={5} visibleFrom="xs">
					<Image
						className={classes.logo}
						src={logo}
						radius="lg"
						h={100}
						w={100}
					/>
					<a
						key="home"
						href="/"
						className={
							homePath ? `${classes.active} ${classes.link}` : `${classes.link}`
						}
					>
						Home
					</a>
					<a
						key="accounts"
						href="/accounts"
						className={
							accountsPath
								? `${classes.active} ${classes.link}`
								: `${classes.link}`
						}
					>
						Accounts
					</a>
				</Group>
			</Container>
		</header>
	);
}
