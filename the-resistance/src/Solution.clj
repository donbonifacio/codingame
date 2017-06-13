(ns Solution
  (:gen-class))

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

(defn starts-with? [s needle]
  (.startsWith (.toString s) needle))

(defn possible-starting-words
  "Given a morse sequence string, and a morse encoded collection of words,
  returns all words that may start the given sequence."
  [morse-sequence morse-words]
  (reduce (fn [starting-words curr]
            (if (starts-with? morse-sequence curr)
              (conj starting-words curr)
              starting-words))
          []
          morse-words))

(defn remove-starting-word
  [morse-sequence word]
  (clojure.string/replace-first morse-sequence
                               (re-pattern word)
                               ""))

(defn possible-word-sequences
  ([morse-sequence morse-words]
   (possible-word-sequences morse-sequence morse-words false))
  ([morse-sequence morse-words verbose]
   (let [starting-words (possible-starting-words morse-sequence morse-words)]
     (when verbose
       (swap! branch-counter inc)
       #_(println "Sequence: " morse-sequence)
       (println "Starting words: " (count starting-words) "/" @branch-counter))
     (->> starting-words
          (map (fn [word]
                  (let [new-morse-sequence (remove-starting-word morse-sequence word)
                        child-messages (possible-word-sequences new-morse-sequence morse-words verbose)]
                    (when verbose
                      #_(println word " -> " new-morse-sequence " : " child-messages))
                    (cond
                      ;; no childs down the line that match
                      (and (zero? child-messages) (not (empty? new-morse-sequence))) 0
                      ;; has a complete child tree
                      (zero? child-messages) 1
                      ;; has several complete child trees
                      :else child-messages)
                      )))
         (reduce +)))))

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
    (println (possible-word-sequences
               morse-sequence
               morse-dictionary true))))
