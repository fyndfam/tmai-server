db.createUser({
  user: "fynd",
  pwd: "password",
  roles: [
    {
      role: "readWrite",
      db: "fynd",
    },
  ],
});
