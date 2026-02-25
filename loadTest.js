import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 100 },
    { duration: '30s', target: 300 },
    { duration: '30s', target: 600 },
    { duration: '30s', target: 0 },
  ],
};

const BASE_URL = 'http://localhost:8080';

export default function () {

  // -----------------------------
  // 1️⃣ CREATE ORDER
  // -----------------------------
  const createPayload = JSON.stringify({
    product_name: "Laptop",
    price: 100,
  });

  const createRes = http.post(
    `${BASE_URL}/create-order`,
    createPayload,
    { headers: { 'Content-Type': 'application/json' } }
  );

  check(createRes, {
    'order created': (r) => r.status === 200,
  });

  const body = JSON.parse(createRes.body);
  const orderId = body.data.ID; // adjust if your JSON differs

  sleep(1);

  // -----------------------------
  // 2️⃣ PAY ORDER
  // -----------------------------
  const payRes = http.patch(`${BASE_URL}/pay/${orderId}`);

  check(payRes, {
    'order paid': (r) => r.status === 200,
  });

  sleep(1);

  // -----------------------------
  // 3️⃣ GET ORDER
  // -----------------------------
  const getRes = http.get(`${BASE_URL}/status/${orderId}`);

  check(getRes, {
    'order fetched': (r) => r.status === 200,
  });

  sleep(1);
}