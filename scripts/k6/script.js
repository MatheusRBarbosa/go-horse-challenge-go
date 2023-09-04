import http from "k6/http";
import { check, sleep } from "k6";
import * as chance from "./chance.min.js";
import { randomLanguage, randomStack } from "./languages.min.js";

const c = chance.Chance();

function makeUser() {
  const person = {
    nickname: `${c.first()}_${c.word({ length: 10 })}`.toLowerCase(),
    name: c.name(),
    birthDate: c.birthday(),
    stack: randomStack(c),
  };

  return person;
}

function makeTerm() {
  const factories = [
    () => c.first().toLowerCase(),
    () => c.name(),
    () => randomLanguage(c),
  ];

  const randomFactory =
    factories[c.integer({ min: 0, max: factories.length - 1 })];

  return randomFactory();
}

export const options = {
  ext: {
    loadimpact: {
      projectID: 3656055,
      name: "GHC: barbosa",
    },
  },
  stages: [{ duration: "5s", target: 1 }],
};

export default function () {
  const baseUrl = "http://localhost:9999";

  let res1 = http.get(`${baseUrl}/api/users`, {
    verb: "get",
    tags: { name: "Get all" },
    headers: {
      "Content-Type": "application/json",
    },
  });
  check(res1, {
    "status is 200": (r) => r.status === 200,
  });

  sleep(1);

  let user = makeUser(); // Some times is generating null to stack
  console.log(user);

  let res2 = http.post(`${baseUrl}/api/users`, JSON.stringify(user), {
    verb: "post",
    tags: { name: "Criação" },
    headers: {
      "Content-Type": "application/json",
    },
  });
  check(res2, {
    "status is 201": (r) => r.status === 201,
  });

  //   let res2 = http.get(`${baseUrl}${res1.headers["Location"]}`, {
  //     verb: "get",
  //     tags: { name: "Consulta" },
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //   });
  //   check(res2, {
  //     "status is 200": (r) => r.status === 200,
  //   });

  //   sleep(1);

  //   let randomTerm = makeTerm();

  //   let res3 = http.get(`${baseUrl}/pessoas?t=${randomTerm}`, {
  //     verb: "get",
  //     tags: { name: "Busca" },
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //   });
  //   check(res3, {
  //     "status is 200": (r) => r.status === 200,
  //   });
}
