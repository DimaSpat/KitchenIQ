import { component$, useSignal, useVisibleTask$ } from '@builder.io/qwik';

export default component$(() => {
  const data = useSignal<string>('');
  const error = useSignal<string>('');

  useVisibleTask$(async () => {
    try {
      const response = await fetch('http://localhost:5000/api/');
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      const text = await response.text();
      data.value = text;
    } catch (e) {
      console.error(e);
      error.value = 'Failed to fetch data from backend.';
    }
  });

  return (
    <>
      <h1>Backend Response:</h1>
      {data.value ? (
        <p>{data.value}</p>
      ) : (
        <p>{error.value || 'Loading...'}</p>
      )}
    </>
  );
});