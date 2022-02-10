import React from 'react';
import { useState } from 'react';
import {NotificationContainer} from 'react-notifications';
import 'react-notifications/lib/notifications.css';
import HeaderBar from '../components/HeaderBar';
import GoShortPage from './GoShortPage';
import GoLongPage from './GoLongPage';
import GoCustomPage from './GoCustomPage';

function MainPage() {

    const [currentPage, setCurrentPage] = useState("Go Short")

    const displayPage = currentPage => {
        switch (currentPage) {
            case "Go Short":
                return <GoShortPage/>
            case "Go Long":
                return <GoLongPage/>
            case "Go Custom":
                return <GoCustomPage/>
            default:
                return <GoShortPage/>
        }
    }

  return (
    <React.Fragment>
        <HeaderBar handlePageChange={setCurrentPage}/>
        <NotificationContainer/>
        {displayPage(currentPage)}
    </React.Fragment>
  );
}

export default MainPage;