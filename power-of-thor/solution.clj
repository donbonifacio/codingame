(ns Player
  (:gen-class))

(def dirs [{:name "N" :offset [0 -1]}
           {:name "NE" :offset [1 -1]}
           {:name "E" :offset [1 0]}
           {:name "SE" :offset [1 1]}
           {:name "S" :offset [0 1]}
           {:name "SW" :offset [-1 1]}
           {:name "W" :offset [-1 0]}
           {:name "NW" :offset [-1 -1]}])

(defn- abs [a] (if (< a 0) (- a) a))
(def diff (comp abs -))

(defn- add-next-pos
  [dir x y target-x target-y]
  (let [[ox oy] (:offset dir)
        next-x (+ x  ox)
        next-y (+ y oy)]
    (assoc dir :next-pos [next-x next-y]
           :cost (+ (diff target-x next-x) (diff target-y next-y)))))

(defn- lucky-dir
  [x y target-x target-y]
  (->> dirs
       (map #(add-next-pos % x y target-x target-y))
       (sort-by :cost)
       first))

(defn -main [& args]
  (let [lightX (read) lightY (read) initialTX (read) initialTY (read)]
    (while true
      (let [remainingTurns (read)] 
        (loop [curr-x initialTX
               curr-y initialTY]
          (let [lucky (lucky-dir curr-x curr-y lightX lightY)
                [next-x next-y] (:next-pos lucky)]
            (println (:name lucky))
            (recur next-x next-y)))))))




