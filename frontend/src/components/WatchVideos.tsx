import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { WatchVideoInterface } from "../models/IWatchVideo";
import { format } from 'date-fns'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function WatchVideos() {
  const classes = useStyles();
  const [watchVideos, setWatchVideos] = useState<WatchVideoInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getWatchVideos = async () => {
    fetch(`${apiUrl}/watch_videos`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setWatchVideos(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getWatchVideos();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              ข้อมูลการเข้าชมวีดีโอ
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/watch_video/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="10%">
                  ลำดับ
                </TableCell>
                <TableCell align="center" width="20%">
                  วีดีโอ
                </TableCell>
                <TableCell align="center" width="20%">
                  ความละอียด
                </TableCell>
                <TableCell align="center" width="20%">
                  เพลย์ลิสต์
                </TableCell>
                <TableCell align="center" width="30%">
                  วันที่และเวลา
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {watchVideos.map((item: WatchVideoInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Video.Name}</TableCell>
                  <TableCell align="center">{item.Resolution.Value}</TableCell>
                  <TableCell align="center">{item.Playlist.Title}</TableCell>
                  <TableCell align="center">{format((new Date(item.WatchedTime)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default WatchVideos;
