import TestImageRoute from './components/TestImageRoute';

export default function App() {
  return (
    <div className="flex flex-col justify-center items-center">
      <p className="mt-5 text-6xl font-bold text-blue-600">
        Sonic Stream Media Conversion Tool
      </p>

      <TestImageRoute className="mt-10" />
    </div>
  );
}
