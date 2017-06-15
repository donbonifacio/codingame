(ns Solution
  (:gen-class))

(set! *warn-on-reflection* true)
(set! *unchecked-math* true)

(def branch-counter (atom 0))

(defmacro time-info
  "This is a copy of the time macro of std clojure. It adds an addition
  string label and outputs it."
  [info expr]
  `(let [start# (. System (nanoTime))
         ret# ~expr]
     (println (str ~info ": " (/ (double (- (. System (nanoTime)) start#)) 1000000.0)))
     ret#))

(def morse-table {\A ".-" \B "-..." \C "-.-." \D "-.."
                  \E "." \F "..-." \G "--." \H  "...."
                  \I ".." \J ".---" \K "-.-" \L ".-.."
                  \M "--" \N "-." \O "---" \P ".--."
                  \Q "--.-" \R ".-." \S "..." \T "-"
                  \U "..-" \V "...-" \W ".--" \X "-..-"
                  \Y "-.--" \Z "--.."})

(defn morse
  ([text]
   (morse text ""))
  ([text sep]
   (clojure.string/join sep (map #(get morse-table %) (str text)))))

(defn log
  [msg]
  (binding [*out* *err*]
    (println msg)))

(defn word-match [morse-sequence word idx]
  (loop [word-idx 0
         morse-idx idx]
    (cond
      (>= word-idx (count word))
        (+ idx word-idx)
      (>= morse-idx (count morse-sequence))
        -2
      (not= (nth morse-sequence morse-idx) (nth word word-idx))
        -1
      :else
        (recur (inc word-idx) (inc morse-idx)))))

(defn morse-counter
  ([^String morse-sequence ^long curr-idx dictionary]
   (loop [words dictionary
         counter 0]
     (if (not (seq words))
       counter
       (let [word (first words)
             ^long next-idx (word-match morse-sequence word curr-idx)]
         (cond
           (= next-idx (count morse-sequence))
            (do
              (println (swap! branch-counter inc))
             1)
 
           (neg? next-idx)
             (recur (rest words) counter)
 
           :else
             (recur (rest words) (+ counter (morse-counter morse-sequence next-idx dictionary)))))))))

(defn load-dictionary-from-stdin
  [dictionary-size]
  (loop [i dictionary-size
         words []]
    (if (= i 0)
      words
      (let [W (read)]
        (recur (dec i) (conj words W))))))

(defn -main [& args]
  (let [morse-sequence (read)
        n-words (read)
        dictionary (time-info "Load dictionary" (load-dictionary-from-stdin n-words))
        morse-dictionary (time-info "Convert dictionary to morse" (doall (map morse dictionary)))]
    (log morse-sequence)
    (log (str "Dictionary words: " (count morse-dictionary)))
    (println (morse-counter
               (str morse-sequence)
               0
               morse-dictionary))))
