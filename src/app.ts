import express, { Request, Response } from "express";

const app = express();

app.get("/", (req: Request, res: Response) => {
  res.status(200).json("delivered");
});

const PORT = 8080;
app.listen(PORT, () => {
  console.log(`Server running on PORT=${PORT}`);
});
