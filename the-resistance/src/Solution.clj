(ns Solution
  (:gen-class))

(def morse-table
  "Morse table used on the challenge."
  {\A ".-" \B "-..." \C "-.-." \D "-.."
   \E "." \F "..-." \G "--." \H  "...."
   \I ".." \J ".---" \K "-.-" \L ".-.."
   \M "--" \N "-." \O "---" \P ".--."
   \Q "--.-" \R ".-." \S "..." \T "-"
   \U "..-" \V "...-" \W ".--" \X "-..-"
   \Y "-.--" \Z "--.."})

(defn morse
  "Given a raw uppercase text, returns it encoded as morse code. It may
  receive a separator."
  ([text]
   (morse text ""))
  ([text sep]
   (clojure.string/join sep (map #(get morse-table %) (str text)))))

(def idx-cache
  "Sores work that has been done before and may be used again. Mainly,
  it will store string indexes and the number of messages originated
  from that point."
  (atom {}))

(defn clear-cache
  "Utility that clears the cache. Used for unit tests."
  []
  (reset! idx-cache {}))

(declare with-indexed-morse-counter)

(defn _with-indexed-morse-counter
  "Worker funtion. Will iterate the given morse sequence, starting at
  curr-idx and will match words from the given dictionary. Returns
  the number of different possible messages from curr-idx."
  ([^String morse-sequence ^long curr-idx dictionary]
   (loop [counter 0
          curr-char-idx curr-idx
          index []]
     (if (= (count morse-sequence) curr-char-idx)
       counter ;; morse-sequence finished, return counter
       (let [curr-symbol (nth morse-sequence curr-char-idx)
             index (conj index curr-symbol)
             match (get-in dictionary index)]
         ;; at this point we have the current symbol being analysed and we
         ;; build the current index with it. Then we try to match it with a
         ;; word on the dictionary
         (cond

           ;; if we're at the last morse sequence char and we have a match
           ;; return the current counter plus the number of possible words
           ;; for the given match
           (and (= curr-char-idx (dec (count morse-sequence)))
                (:word? match))
             (+ counter (:matches match))

           ;; if we don't have a match means that no ditionary words are
           ;; possible at this point. Return the current counter
           (nil? match)
             counter

           ;; the current index matches a word, let's branch this and start
           ;; counting again from the next position, but from a fresh
           ;; dictionary view
           (:word? match)
             (recur (+ counter 
                       (* (:matches match)
                          (with-indexed-morse-counter morse-sequence
                                                      (inc curr-char-idx)
                                                      dictionary)))
                    (inc curr-char-idx)
                    index)

           ;; we got a match, but no word, proceed to next char to see
           ;; if we find a match on the next positions
           :else
             (recur counter (inc curr-char-idx) index)))))))

(defn with-indexed-morse-counter
  "Wraps the main worker function to memoize it. It will store the result for
  a given curr-idx and next calls will return it rigt away. This could be
  a memoize and all it does more things because they were useful for testing."
  [^String morse-sequence ^long curr-idx dictionary]
  (let [idx-counter (get @idx-cache curr-idx)
        counter (:counter idx-counter)
        visits (inc (or (:visits idx-counter) 0))
        value (or counter (_with-indexed-morse-counter morse-sequence curr-idx dictionary))]
      (swap! idx-cache assoc curr-idx {:counter value :visits visits})
      value))

(defn index-dictionary
  "Creates an index from the given words. It will create a tree and whenever
  we have a word, we'll have a :word? true field. It also stores the number
  of times a given combination is registered (different words may have the
  same spaceless morse sequence.)"
  [words]
  (reduce (fn [indexed word]
            (let [index (vec word)
                  mapper (get-in indexed index)]
              (-> indexed
                  (assoc-in (conj (vec word) :matches) (inc (or (:matches mapper) 0)))
                  (assoc-in (conj (vec word) :word?) true))))
          {}
          words))

(defn indexed-morse-counter
  "Wraps the worker function to convert the word dictionary in a indexed tree."
  [^String morse-sequence ^long curr-idx dictionary]
  (with-indexed-morse-counter morse-sequence curr-idx
    (index-dictionary dictionary)))

(defn load-dictionary-from-stdin
  "Just returns all words from the stdin."
  [dictionary-size]
  (loop [i dictionary-size
         words []]
    (if (= i 0)
      words
      (let [W (read)]
        (recur (dec i) (conj words W))))))

(defn -main
  "The main function, will read data from the stdin and call the worker
  function."
  [& args]
  (let [morse-sequence (read)
        n-words (read)
        dictionary (load-dictionary-from-stdin n-words)
        morse-dictionary (map morse dictionary)]
    (println (indexed-morse-counter
               (str morse-sequence)
               0
               morse-dictionary))))
