import { Button } from "@material-ui/core";
import { MDBCol } from "mdbreact";

function Search() {
  return (
    <MDBCol style={{flexDirection:'row'}} md="2">
        <input className="form-control" type="text" placeholder="Search" aria-label="Search" />
        <Button size="sm" variant="outlined" color="primary">Search</Button>
    </MDBCol>
  );
}

export default Search;