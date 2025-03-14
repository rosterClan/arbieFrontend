import React , {useState, useEffect, useContext} from 'react';
import { useLocation } from 'react-router-dom';

import NewOddsTable from './components/new_odds_table/new_odds_table';
import RaceHeader from './components/race_banner/race_view';
import RelatedRace from './components/related_races/related_races';
import EntrantGraphs from '../analytics/components/graphs/graphs';

import './race_view.css'

const RaceView = (props) => {
  const location = useLocation();
  const queryParams = new URLSearchParams(location.search);

  const track = queryParams.get('track');
  const round = queryParams.get('round');
  const startTime = queryParams.get('start_time');

  const [race_data, set_race_data] = useState({});

  useEffect(()=>{
    fetch(`http://127.0.0.1:8080/get_race_details?track=${track}&round=${round}&start_time=${startTime}`)
    .then(data => data.json())
    .then(data => set_race_data(data))
    .catch(err => console.log(err))
  },[]);

  return (
    <>
      <RaceHeader track_name={`${track} R${round}`}/>
      <NewOddsTable entrants={race_data['Entrants']} />
      <EntrantGraphs entrants={race_data['Entrants']} />
    </>
  );
}

export default RaceView; 
