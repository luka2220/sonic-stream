interface PropTypes {
  className: string;
}

export default function TestImageRoute({ className }: PropTypes) {
  return (
    <div className={`text-3xl text-amber-300 ${className}`}>
      <p>Form for backend image route testing</p>
    </div>
  );
}
