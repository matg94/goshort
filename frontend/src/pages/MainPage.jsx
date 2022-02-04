import { AppBar, Button, Paper } from '@material-ui/core';
import React from 'react';
import { useState } from 'react';
import { makeStyles } from '@material-ui/core';
import HeaderBar from './HeaderBar';
import GoShortPage from './GoShortPage';
import GoLongPage from './GoLongPage';

function MainPage() {

    const [currentPage, setCurrentPage] = useState("Go Short")

    const displayPage = currentPage => {
        switch (currentPage) {
            case "Go Short":
                return <GoShortPage/>
            case "Go Long":
                return <GoLongPage/>
            case "Go Custom":
                return <h1>CUSTOM NOT YET IMPL</h1>
        }
    }

  return (
    <React.Fragment>
        <HeaderBar handlePageChange={setCurrentPage}/>
        {displayPage(currentPage)}
    </React.Fragment>
  );
}

export default MainPage;