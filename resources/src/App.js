import React from 'react';
import Header from './components/Header';
import Footer from './components/Footer';
import Home from './pages/Home';
import About from './pages/About';
import './styles/App.css';

function App({ page }) {
    let PageComponent;

    switch(page) {
        case 'home':
            PageComponent = Home;
            break;
        case 'about':
            PageComponent = About;
            break;
        default:
            PageComponent = () => <div>Page not found</div>;
    }

    return (
        <div>
            <Header />
            <main>
                <PageComponent />
            </main>
            <Footer />
        </div>
    );
}

export default App;
