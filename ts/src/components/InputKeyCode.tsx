import React, { KeyboardEventHandler, useEffect, useRef, useState } from "react";
import { useAsyncFn } from "react-use";
import { getKeyCode } from "../gobind";
import { getKeyFromVKCode } from "../util/keycode";

interface Props {
  keyCode: number;
  onChange: (c: number) => void;
}

export default function InputKeyCode(props: Props): JSX.Element {
  const [text, setText] = useState(getKeyFromVKCode(props.keyCode) ?? `${props.keyCode}`);
  const [state, doRequest] = useAsyncFn(getKeyCode);
  const [isSent, setIsSent] = useState(false);
  const ref = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (!isSent) {
      return;
    }

    if (state.loading || !state.value) {
      setText("...");
      return;
    }

    props.onChange(state.value);
    ref.current?.blur();
    const key = getKeyFromVKCode(state.value);
    if (key) {
      setText(key);
    }
  }, [state.loading]);

  const request = () => {
    setIsSent(true);
    doRequest();
  };

  const onKeyDown: KeyboardEventHandler<HTMLInputElement> = (e) => {
    if (state.loading) {
      setText(e.key);
    }
  };

  const onChange = (c: number) => {
    props.onChange(c);
  };

  return (
    <div
      style={{
        display: "flex",
        flexDirection: "row",
      }}
    >
      <input ref={ref} autoFocus readOnly value={text} size={14} onFocus={request} onKeyDown={onKeyDown} />
      <button
        onClick={() => {
          onChange(0);
        }}
      >
        X
      </button>
    </div>
  );
}
