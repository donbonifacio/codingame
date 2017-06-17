(ns Solution
  (:gen-class))

#_(set! *warn-on-reflection* true)
#_(set! *unchecked-math* true)

(def branch-counter (atom 0))

(defmacro time-info
  "This is a copy of the time macro of std clojure. It adds an addition
  string label and outputs it."
  [info expr]
  `(let [start# (. System (nanoTime))
         ret# ~expr]
     #_(println (str ~info ": " (/ (double (- (. System (nanoTime)) start#)) 1000000.0)))
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
  #_(binding [*out* *err*]
    (println msg)))

(defn word-match [^String morse-sequence word ^long idx]
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
              #_(println (swap! branch-counter inc))
             1)
 
           (neg? next-idx)
             (recur (rest words) counter)
 
           :else
             (recur (rest words) (+ counter (morse-counter morse-sequence next-idx dictionary)))))))))

(def cache (atom {}))

(declare with-indexed-morse-counter)

(defn _with-indexed-morse-counter
  ([^String morse-sequence ^long curr-idx dictionary]
   (loop [counter 0
          curr-char-idx curr-idx
          index []]
     (if (= (count morse-sequence) curr-char-idx)
       counter
       (let [curr-symbol (nth morse-sequence curr-char-idx)
             index (conj index curr-symbol)
             match (get-in dictionary index)]
         (cond
           (and (= curr-char-idx (dec (count morse-sequence)))
                (:word? match))
             (:matches match)
           (nil? match)
             counter
           (:word? match)
            (recur (+' counter (* (:matches match) (with-indexed-morse-counter morse-sequence
                                                          (+ 1 curr-char-idx)
                                                          dictionary)))
                   (+ 1 curr-char-idx)
                   index)
          :else
            (recur counter (inc curr-char-idx) index)))
       ))))

(def idx-cache (atom {}))

(defn with-indexed-morse-counter
  [^String morse-sequence ^long curr-idx dictionary]
  (let [idx-counter (get @idx-cache curr-idx)
        counter (:counter idx-counter)
        visits (inc (or (:visits idx-counter) 0))
        value (or counter (_with-indexed-morse-counter morse-sequence curr-idx dictionary))]
      #_(prn "Generated for idx" curr-idx "counter" value "visits" visits)
      (swap! idx-cache assoc curr-idx {:counter value :visits visits})
      value)
      )

(defn clear-cache [] (reset! idx-cache {}))

(defn index-dictionary [words]
  (reduce (fn [indexed word]
            (let [index (vec word)
                  mapper (get-in indexed index)]
              (-> indexed
                  (assoc-in (conj (vec word) :matches) (inc (or (:matches mapper) 0)))
                  (assoc-in (conj (vec word) :word?) true))))
          {}
          words))

(defn indexed-morse-counter
  [^String morse-sequence ^long curr-idx dictionary]
  (with-indexed-morse-counter morse-sequence curr-idx
    (index-dictionary dictionary)))

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
        morse-dictionary (time-info "Convert dictionary to morse" (map morse dictionary))]
    (log (str "String size:" (count (str morse-sequence))))
    (log (str "Dictionary words: " (count morse-dictionary)))
    (println (indexed-morse-counter
               (str morse-sequence)
               0
               morse-dictionary))))
