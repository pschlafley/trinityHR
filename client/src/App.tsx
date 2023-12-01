/* eslint-disable no-mixed-spaces-and-tabs */
import '@mantine/core/styles.css';
import { HeroComponent } from './components/heroComponent/HeroComponent';
import { NavBar } from './components/Navbar/Navbar';
import { Outlet } from 'react-router-dom';

function App() {
	return (
		<div>
			<NavBar />
			<HeroComponent />

			<Outlet />
		</div>
	);
}

export default App;
