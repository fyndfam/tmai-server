db.createUser({
  user: "tmai",
  pwd: "password",
  roles: [
    {
      role: "readWrite",
      db: "tmai",
    },
  ],
});
