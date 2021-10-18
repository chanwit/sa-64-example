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
import { UsersInterface } from "../models/IUser";

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

function Users() {
  const classes = useStyles();
  const [users, setUsers] = useState<UsersInterface[]>([]);

  const getUsers = async () => {
    const apiUrl = "http://localhost:8080/users";
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setUsers(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getUsers();
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
              ข้อมูลสมาชิก
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/user/create"
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
                <TableCell align="center" width="45%">
                  ชื่อ
                </TableCell>
                <TableCell align="center" width="45%">
                  อีเมล
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {users.map((user: UsersInterface) => (
                <TableRow key={user.ID}>
                  <TableCell align="center">{user.ID}</TableCell>
                  <TableCell align="center">{user.Name}</TableCell>
                  <TableCell align="center">{user.Email}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Users;
