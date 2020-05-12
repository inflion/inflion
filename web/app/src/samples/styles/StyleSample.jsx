import { makeStyles, Paper } from '@material-ui/core';
import React from 'react';

// スタイルの記述をする
const useStyles = makeStyles((theme) => ({
  outer: {
    width: '100%',
    height: '100%',
    padding: theme.spacing(4),
    backgroundColor: theme.palette.background.default,
  },

  inner: {
    width: '720px',
    minHeight: '420px',

    margin: '0 auto',
    padding: theme.spacing(2),

    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
}));

export const StyleSample = (props) => {
  // useStyles() を呼ぶと上記スタイリングが適応されるユニークなクラスネームが取得できる
  const classes = useStyles();

  return (
    //   各コンポーネントにスタイルをあてる
    <div className={classes.outer}>
      <Paper className={classes.inner}>Hooks でクラススタイルが書きやすくなったよ</Paper>
    </div>
  );
};
